package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/usecase"
	"net/http"
)

type loginInput struct {
	Email    string `json:email`
	Password string `json:password`
}

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var input loginInput
	if bindError := c.BindJSON(&input); bindError != nil {
		fmt.Println("Error Occur")
		c.Status(http.StatusBadRequest)
	}

	status, err := handler.authUsecase.Login(input.Email)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if status <= 0 {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	token, _ := handler.authUsecase.GenToken(input.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func NewAuthHandler(usecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase}
}
