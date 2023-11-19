package main

import (
	"github.com/gin-gonic/gin"
	"violin-home.cn/retail/common"
	"violin-home.cn/retail/config"
	"violin-home.cn/retail/router"
	"violin-home.cn/retail/store"

	// activate router
	_ "violin-home.cn/retail/api/v1"
)

func main() {

	r := gin.Default()

	router.InitRouter(r)

	store.NewRedisClient()
	store.NewMongoClient()

	common.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)

}
