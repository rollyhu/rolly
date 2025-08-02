package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/test/init_project/config"
	"github.com/test/init_project/form"
)

// 实现文章的创建功能
func CreateComment(c *gin.Context) {
	userId := c.MustGet("userID").(int)
	// config.MyLoger.Debug().Uint("userID", c.MustGet("userID").(uint)).Msg("    CreatePost")
	if userId == 0 {
		c.JSON(404, gin.H{"msg": "Please authenticate!"})
		return
	}
	var comment form.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(404, gin.H{"msg": "request post err"})
		return
	}
	// cc := c.MustGet("userID").(uint)
	// fmt.Println("MustGet     ", cc)
	erro := config.Db.Create(&comment).Error
	if erro != nil {
		config.MyLoger.Debug().Msg(erro.Error())
		c.JSON(404, gin.H{"msg": " comment create fail"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": " comment create successfully"})
		return
	}
}

// 实现文章的读取功能
func GetComments(c *gin.Context) {
	config.MyLoger.Debug().Msg("GetComments--->")
	iD := c.Param("postID")
	fmt.Println("postID", iD)
	// config.MyLoger.Debug().Int("postID", postID).Send()
	var comments []form.Comment
	err := config.Db.Model(&form.Comment{}).Where("post_id=?", iD).Find(&comments).Error
	if err != nil {
		c.JSON(500, gin.H{"msg": "get comments fail!"})
		config.MyLoger.Error().Str("get err", err.Error())
		return
	}
	c.JSON(http.StatusOK, comments)

}
