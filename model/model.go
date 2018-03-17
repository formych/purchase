package model

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/formych/purchase/dao"

	"github.com/formych/util"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// PurchaseInfo 对应页面的字段
type PurchaseInfo struct {
	User        string `form:"user" json:"user"`
	Company     string `form:"company" json:"company"`
	Tel         string `form:"tel" json:"tel"`
	PurchaseNum int    `form:"purchase_num" json:"purchase_num"`
	PuchaseTime string `form:"purchase_time" json:"purchase_time"`
}

// Index ...
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}

// Add record to db
func Add(c *gin.Context) {
	p := &PurchaseInfo{}
	c.Param("purchase_time")
	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		log.Println(err)
		return
	}
	log.Println(c.Param("purchase_time"))
	tnow := time.Now().Format("2006-01-02 15:04:05")
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
		c.String(http.StatusBadGateway, "添加数据失败")
		return
	}
	c.Redirect(302, "status")
}

// GetExcel download excel file to user
func GetExcel(c *gin.Context) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	var tmpfile = "./excel/" + util.NonceStr(15) + ".xlsx"

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		log.Printf(err.Error())
	}
	res, err := dao.PurchaseInfoDao.Get()
	if err != nil {
		log.Println("DB get records failed")
		return
	}
	row = sheet.AddRow()
	row.WriteSlice(&[]string{"姓名", "公司", "电话", "采购数量", "采购时间"}, 5)
	for _, v := range res {
		row = sheet.AddRow()
		row.WriteStruct(&PurchaseInfo{v.User, v.Company, v.Tel, v.PurchaseNum, v.PurchaseTime}, 5)
	}

	if err = file.Save(tmpfile); err != nil {
		log.Println(err.Error())
	}
	defer os.Remove(tmpfile)

	f, err := os.Open(tmpfile)
	defer f.Close()
	if err != nil {
		log.Println(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	c.Header("Content-Disposition", "attachment; filename=purchaseinfo.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}

// List show all records
func List(c *gin.Context) {
	// exeSQL := "SELECT * FROM %s "
	res, err := dao.PurchaseInfoDao.Get()
	if err != nil {
		log.Printf("DB get record failed, err: [%v]", err)
		c.String(http.StatusBadGateway, "获取数据失败")
		return
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"Rows": res,
	})
}

// Status 默认成功页面
func Status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", gin.H{})
}
