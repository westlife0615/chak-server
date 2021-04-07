package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
