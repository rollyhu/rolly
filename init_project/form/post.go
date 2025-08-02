package form

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null";json:"Title" binding:"required`
	Content string `gorm:"not null";json:"Content" binding:"required`
	UserID  uint   `json:"userId"`
	User    User
}
