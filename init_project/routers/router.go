package routers

import "github.com/gin-gonic/gin"

func SetRouter(router *gin.Engine) {
	UserGroup := router.Group("/u")
	PostGroup := router.Group("/p")
	CommentGroup := router.Group("/c")
	UserGroup.GET("user")
	PostGroup.GET("")
	CommentGroup.GET("")
}
