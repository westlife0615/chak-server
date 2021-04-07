package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/westlife0615/chak-server/service"
	"net/http"
)

//StartGin function
func StartGin() {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/login", (&service.AuthService{}).Register)
		auth.POST("/register", (&service.AuthService{}).Register)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}