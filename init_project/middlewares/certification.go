package middlewares

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/handlers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查如果是GET 则跳过验证
		config.MyLoger.Debug().Str("ctx.Request.Method", ctx.Request.Method).Str(" ctx.FullPath()", ctx.FullPath()).Send()
		if ctx.Request.Method == "GET" {
			ctx.Next()
			return
		}

		header := ctx.GetHeader("x-Token")
		if header == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "缺少认证"})
			ctx.Abort()
			return
		}
		config.MyLoger.Debug().Str("header", header).Send()
		mytonken, err := jwt.ParseWithClaims(header, &handlers.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})
		config.MyLoger.Debug().Str("header", strconv.Itoa(mytonken.Claims.(*handlers.MyClaims).UserId)).Send()
		ctx.Set("userID", mytonken.Claims.(*handlers.MyClaims).UserId)
		if err != nil {
			config.MyLoger.Debug().Str("err", err.Error()).Send()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "认证失效"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
