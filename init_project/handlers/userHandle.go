package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/form"
	"golang.org/x/crypto/bcrypt"
)

// // 注册请求体
// type RegisterRequest struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// 	Email    string `json:"email" binding:"required,email"`
// }

// // 登录请求体
// type LoginRequest struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func Register(c *gin.Context) {
	fmt.Println("Register running!")
	var user form.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	// var user1 form.User
	// user1.Email = user.Email
	// user1.Username = user.Username
	// user1.Password = user.Password
	if err := config.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

type MyClaims struct {
	UserId int `json:"userId`
	jwt.StandardClaims
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	fmt.Println("Login......")
	var user LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user1 form.User
	// if err := config.Db.Model(&form.User{}).Where("username = ?", user.Username).First(&user1).Error; err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	// 	return
	// }
	err := config.Db.Model(&form.User{}).Where("username = ?", user.Username).First(&user1).Error
	config.MyLoger.Debug().Str("userName", user1.Username).Send()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username！"})
		return
	}

	// 验证密码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	config.MyLoger.Debug().Str("passWord", string(hashedPassword)).Send()
	config.MyLoger.Debug().Str("passWordByGorm", user1.Password).Send()
	if err := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		UserId: int(user1.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 5000,
			Issuer:    "rolly",
			NotBefore: time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	//打印token
	// fmt.Println("tokenString:", tokenString)

	c.JSON(http.StatusOK, gin.H{"x-Token": tokenString, "msg": user1.Username + "  login successfully"})

	// 剩下的逻辑...
}
