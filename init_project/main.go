package main

import (
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/routers"
)

func main() {

	config.InitDB()
	router := gin.Default()
	routers.SetRouter(router)

}
