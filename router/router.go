package router

import (
	"github.com/formych/purchase/model"
	"github.com/gin-gonic/gin"
)

// Router 全局路由注册
var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Static("/assets", "./assets")
	Router.LoadHTMLGlob("views/*")

	Router.GET("/", model.Index)
	Router.GET("/excel", model.GetExcel)
	Router.POST("/add", model.Add)
}
