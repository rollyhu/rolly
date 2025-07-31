package config

import (
	"fmt"

	"github.com/test/init_project/form"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// 初始gormDB
func InitDB() {

	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	Db = db.Debug()
	initTables(Db)

}

// 初始表
func initTables(db *gorm.DB) {
	err := db.AutoMigrate(&form.User{}, &form.Post{}, &form.Comment{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
