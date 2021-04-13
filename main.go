package main

import (
	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/handler"
	"github.com/westlife0615/chak-server/usecase"
	"github.com/westlife0615/chak-server/repository"
)

func main() {
	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	r := gin.Default()
	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PATCH("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.Run(":8080")
}