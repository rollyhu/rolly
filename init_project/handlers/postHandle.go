package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/form"
)

// 实现文章的创建功能
func CreatePost(c *gin.Context) {
	userId := c.MustGet("userID").(int)
	// config.MyLoger.Debug().Uint("userID", c.MustGet("userID").(uint)).Msg("    CreatePost")
	var post form.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(404, gin.H{"msg": "request post err"})
		return
	}
	// cc := c.MustGet("userID").(uint)
	// fmt.Println("MustGet     ", cc)
	post.UserID = uint(userId)
	config.MyLoger.Debug().Int("userID", userId).Msg("    CreatePost")
	config.MyLoger.Debug().Str("requstTitle", post.Title).Send()
	erro := config.Db.Create(&post).Error
	if erro != nil {
		config.MyLoger.Debug().Msg(erro.Error())
		c.JSON(404, gin.H{"msg": " post create fail"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": " post create successfully"})
		return
	}
}

// 实现文章的读取功能
func GetPosts(c *gin.Context) {
	config.MyLoger.Debug().Msg("GetPosts--->")
	var posts []form.Post
	err := config.Db.Find(&posts).Error
	if err != nil {
		c.JSON(500, gin.H{"msg": " getPost fail!"})
		return
	}
	c.JSON(http.StatusOK, posts)

}

// 实现单个文章的读取
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	config.MyLoger.Debug().Msg("GetPostById--->" + id)
	var post form.Post
	err := config.Db.Preload("User").First(&post, id).Error
	if err != nil {
		config.MyLoger.Error().Str("err", err.Error()).Send()
		c.JSON(500, gin.H{"msg": " getPostById fail!"})
		return
	}
	post.User.Password = ""
	c.JSON(http.StatusOK, post)
}

// 实现文章的更新功能
func UpdatePost(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	// config.MyLoger.Debug().Uint("userID", c.MustGet("userID").(uint)).Msg("    CreatePost")
	var post form.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		config.MyLoger.Error().Str("err", err.Error()).Send()
		c.JSON(400, gin.H{"msg": "request fail"})
		return
	}
	if userID != int(post.UserID) {
		c.JSON(403, gin.H{"msg": "No permission to update"})
		return
	}
	post.CreatedAt = time.Now()
	if err := config.Db.Save(&post).Error; err != nil {
		config.MyLoger.Error().Str("err", err.Error()).Send()
		c.JSON(500, "post update fail")
		return
	}
	c.String(200, "update sucessful!")

}

// 实现文章的删除功能
func DelPost(c *gin.Context) {
	fmt.Println("DelPost---------------->>>>")
	userID := c.MustGet("userID").(int)
	//获取前端传来的userID
	// tmUser_id := c.Param("userId")
	// postId := c.Param("postId")
	var post form.Post
	c.ShouldBindJSON(&post)
	config.MyLoger.Debug().Uint("post.ID", post.ID).Uint("userID", post.UserID).Send()
	if post.UserID != uint(userID) {
		c.JSON(403, gin.H{"msg": "No permission to del"})
		return
	}
	if err := config.Db.Delete(&form.Post{}, post.ID).Error; err != nil {
		c.JSON(500, gin.H{"msg": "post del fail"})
		config.MyLoger.Error().Str("err", err.Error()).Send()
		return
	}
	c.String(200, "del sucessful!")
}
