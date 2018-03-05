package main

import (
	"github.com/formych/purchase/dao"
	"github.com/formych/purchase/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 暂时使用显式关闭数据库连接
	defer dao.DB.Close()

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")

	router.GET("/", model.Index)
	router.GET("/excel", model.GetExcel)
	router.POST("/add", model.Add)

	router.Run(":8080")
}
