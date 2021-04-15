package handler

import (
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
	tu usecase.UserUsecase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input createUserInput
	c.Bind(&input)
	user := model.User{Email: input.Email, Password : input.Password, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := h.tu.CreateUser(user)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input updateUserInput
	c.BindJSON(&input)
	user := model.User{Id: input.Id, Email: input.Email, UpdatedAt: time.Now()}
	err := h.tu.UpdateUser(user)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusBadRequest)
	}

	if err := h.tu.DeleteUser(id); err != nil {
		log.Printf("Error occured -  %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.tu.GetAllUsers()
	c.JSON(200, users)
}

func NewUserHandler(tu usecase.UserUsecase) *UserHandler {
	return &UserHandler{tu}
}
