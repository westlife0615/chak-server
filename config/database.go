package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/westlife0615/chak-server/model"
)

var Database *gorm.DB

func InitDatabase (){
	var err error
	Database, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	migrate()
}

func migrate() {
	Database.AutoMigrate(&model.User{})
}