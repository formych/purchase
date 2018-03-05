package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/formych/purchase/dao"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type PurchaseInfo struct {
	User        string `form:"user" json:"user"`
	Company     string `form:"company" json:"company"`
	Tel         string `form:"tel" json:"tel"`
	PurchaseNum int    `form:"purchase_num" json:"purchase_num"`
	PuchaseTime string `form:"purchase_time" json:"purchase_time"`
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}
func Add(c *gin.Context) {
	p := &PurchaseInfo{}
	c.Param("purchase_time")
	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		log.Println(err)
		return
	}
	fmt.Println(p)
	fmt.Println(c.Param("purchase_time"))
	tnow := time.Now().Format("2016-01-02 15:04:05")
	r := &dao.PurchaseInfo{
		User:         p.User,
		Company:      p.Company,
		Tel:          p.Tel,
		PurchaseNum:  p.PurchaseNum,
		PurchaseTime: p.PuchaseTime,
		CreatedTime:  tnow,
		UpdatedTime:  tnow,
	}
	err := dao.PurchaseInfoDao.Add(r)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetExcel(c *gin.Context) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	var tmpfile = "./excel/tmp.xlsx"

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := dao.PurchaseInfoDao.Get()
	if err != nil {
		log.Println("Get excel record failed")
		return
	}
	row = sheet.AddRow()
	row.WriteSlice(&[]string{"用户名", "公司", "电话", "采购数量", "采购时间"}, 5)
	defer os.Remove(tmpfile)
	for _, v := range res {
		row = sheet.AddRow()
		row.WriteStruct(&PurchaseInfo{v.User, v.Company, v.Tel, v.PurchaseNum, v.PurchaseTime}, 5)
		err = file.Save(tmpfile)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	file1, err := os.Open(tmpfile)
	defer file1.Close()
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file1)
	if err != nil {
		fmt.Printf(err.Error())
	}
	c.Header("Content-Disposition", "attachment; filename=purchaseinfo.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}
