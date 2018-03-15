package main

import (
	"github.com/formych/purchase/dao"
	"github.com/formych/purchase/router"
)

func main() {
	dao.Init()
	router.Router.Run(":8080")
}
