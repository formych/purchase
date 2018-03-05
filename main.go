package main

import (
	"github.com/formych/purchase/dao"
	"github.com/formych/purchase/router"
)

func main() {
	// 暂时使用显式关闭数据库连接
	defer dao.DB.Close()
	router.Router.Run(":8080")
}
