package main

import (
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/routers"
)

func main() {

	config.InitDB()
	config.InitLog()
	router := gin.Default()
	routers.SetRouter(router)
	router.Run(":8080")

}
