package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/config"
	"github.com/westlife0615/chak-server/handler"
	"github.com/westlife0615/chak-server/repository"
	"github.com/westlife0615/chak-server/usecase"
	"net/http"
	"strings"
)

func main() {
	config.InitDatabase()
	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(authUsecase)

	r := gin.Default()
	r.Use(CORSMiddleware())
	//r.Use(BearerAuthMiddleware())
	r.GET("/users", userHandler.GetAllUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users", BearerAuthMiddleware(), userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.POST("/login", authHandler.Login)
	r.Run(":8080")
}

func BearerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "login") {
			c.Next()
			return
		}
		encodedToken := c.Request.Header.Get("Authorization")
		_, err := jwt.Parse(strings.Replace(encodedToken, "bearer ", "", -1), func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				return nil, fmt.Errorf("Invalid token", token.Header["alg"])

			}
			return []byte("westlife0615"), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
