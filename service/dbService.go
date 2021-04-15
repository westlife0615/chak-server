package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/westlife0615/chak-server/model"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate() {
	db := Connect()
	db.AutoMigrate(&model.User{})
}

func seed(db *gorm.DB) {
	users := []model.User{
		//{Email: "test@test.com", Username: "Joe420"},
		//{Email: "yes@yes.com", Username: "Bob"},
	}
	for _, u := range users {
		db.Create(&u)
	}

	var joe, bob model.User
	db.First(&joe, "Username = ?", "Joe420")
	db.First(&bob, "Username = ?", "Bob")

}
