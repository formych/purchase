package main

import (
	"github.com/formych/purchase/dao"
	"github.com/formych/purchase/router"
)

func main() {
	dao.Init()
	defer dao.DB.Close()
	router.Router.Run(":8080")
}
