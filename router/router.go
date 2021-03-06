package router

import (
	"github.com/formych/purchase/api"
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
	Router.GET("/list", model.List)
	Router.GET("/status", model.Status)
	Router.POST("/add", model.Add)

	group := Router.Group("/v1")
	group.GET("/list", api.List)
}
