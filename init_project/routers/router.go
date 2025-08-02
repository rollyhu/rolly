package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/handlers"
	"github.com/test/init_project/middlewares"
)

func SetRouter(router *gin.Engine) {
	//分组
	UserGroup := router.Group("/u")
	PostGroup := router.Group("/p")
	CommentGroup := router.Group("/c")

	//分组执行
	UserGroup.POST("user", func(ctx *gin.Context) { handlers.Register(ctx) })
	UserGroup.POST("login", func(ctx *gin.Context) { handlers.Login(ctx) })

	//post路由
	PostGroup.Use(middlewares.AuthMiddleware())
	PostGroup.POST("post", func(ctx *gin.Context) { handlers.CreatePost(ctx) })
	PostGroup.GET("post", func(ctx *gin.Context) { handlers.GetPosts(ctx) })
	PostGroup.GET("post/:id", func(ctx *gin.Context) { handlers.GetPostById(ctx) })
	PostGroup.PUT("post", func(ctx *gin.Context) { handlers.UpdatePost(ctx) })
	PostGroup.DELETE("post", func(ctx *gin.Context) { handlers.DelPost(ctx) })

	//Comment路由
	CommentGroup.Use(middlewares.AuthMiddleware())
	CommentGroup.POST("conment", func(ctx *gin.Context) { handlers.CreateComment(ctx) })
	CommentGroup.GET("conment/:postID", func(ctx *gin.Context) { handlers.GetComments(ctx) })

}
