package service

import (
	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/model"
	"net/http"
)

//type AuthService interface {
//	login(email string, password string)
//	register(email string, password string)
//}

type AuthService struct {}

func (auth *AuthService) Register(c *gin.Context){
	var data model.User
	db := Connect()
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&data)

	c.JSON(200, gin.H{"message": "created"})
}

