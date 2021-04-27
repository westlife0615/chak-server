package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/model"
	"github.com/westlife0615/chak-server/usecase"
)

type createUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type updateUserInput struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type UserHandler struct {
	usecase usecase.UserUsecase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input createUserInput
	bindError := c.Bind(&input)
	if bindError != nil {
		log.Printf("Error occured - %+v", bindError)
		c.Status(http.StatusInternalServerError)
	}


	user := model.User{Email: input.Email, Password: input.Password, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := h.usecase.CreateUser(user)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input updateUserInput
	bindError := c.BindJSON(&input)
	if bindError != nil {
		log.Printf("error occured %v" , bindError)
		c.Status(http.StatusInternalServerError)
	}

	fmt.Print(input)
	user := model.User{Id: input.Id, Email: input.Email, UpdatedAt: time.Now()}

	err := h.usecase.UpdateUser(user)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, &user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusBadRequest)
	}

	if err := h.usecase.DeleteUser(id); err != nil {
		log.Printf("Error occured -  %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.usecase.GetAllUsers()
	c.JSON(200, users)
}

func NewUserHandler(tu usecase.UserUsecase) *UserHandler {
	return &UserHandler{tu}
}
