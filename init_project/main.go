package main

import (
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/routers"
)

func main() {

	// db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"))
	// if err != nil {
	// 	panic(err)
	// }
	// zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// log.Info().Msg("Hello from Zerolog global logger")
	// log.Error().Msg("zero log  erro")

	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "hello!")
	// })
	// router.Run(":8080")

	// logger := zerolog.New(os.Stdout)
	// even := logger.Info()
	// even.Str("name", "rolly").Int("age", 35).Msg("login!")
	// // even.Msg("login1！")
	// buildInfo, _ := debug.ReadBuildInfo()

	// fmt.Println(buildInfo)
	// cpuProfile := runtime.NumCPU()
	// fmt.Println(cpuProfile)
	config.InitDB()
	config.InitLog()
	router := gin.Default()
	routers.SetRouter(router)
	router.Run(":8080")

}
