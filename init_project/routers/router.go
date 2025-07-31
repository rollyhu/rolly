package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/handle"
)

func SetRouter(router *gin.Engine) {
	//分组
	UserGroup := router.Group("/u")
	PostGroup := router.Group("/p")
	CommentGroup := router.Group("/c")

	//分组执行
	UserGroup.POST("user", func(ctx *gin.Context) { handle.Register(ctx) })
	PostGroup.GET("")
	CommentGroup.GET("")
}
