package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/formych/purchase/dao"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageTotal := c.DefaultQuery("page_total", "20")
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageTotalNum, _ := strconv.ParseUint(pageTotal, 10, 64)
	if pageNum == 0 {
		pageNum = 1
	}
	if pageTotalNum == 0 {
		pageTotalNum = 20
	}

	start := (pageNum - 1) * pageTotalNum
	resultTotal, err := dao.PurchaseInfoDao.Count()
	if err != nil {
		resultTotal = 0
	}
	res, err := dao.PurchaseInfoDao.Get(start, pageTotalNum)
	if err != nil {
		log.Printf("DB get record failed, err: [%v]", err)
		c.String(http.StatusBadGateway, "获取数据失败")
		return
	}
	if len(res) == 0 {
		res = []*dao.PurchaseInfo{}
	}
	c.JSON(http.StatusOK, gin.H{
		"result_total": resultTotal,
		"data":         res,
		"page":         pageNum,
		"page_total":   len(res),
		"status_code":  200,
	})
}
