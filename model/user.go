package model

import "time"

type User struct {
	Id        int `json:id gorm:"primary_key"`
	Password  uint
	Email     string    `json:email,omitempty`
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
}
