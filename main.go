package main

import (
	"github.com/westlife0615/chak-server/routes"
)

func main() {
	//var loginService service.LoginService = service.StaticLoginService()
	//var jwtService service.JWTService = service.JWTAuthService()
	//var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	//service.Connect()

	// 기본 gin 엔진
	//server := gin.New()

	//server.POST("/login", func(ctx *gin.Context) {
	//	token := loginController.Login(ctx)
	//	if token != "" {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"token": token,
	//		})
	//	} else {
	//		ctx.JSON(http.StatusUnauthorized, nil)
	//	}
	//})
	//port := "8080"
	//server.Run(":" + port)

	//service.Migrate()
	routes.StartGin()

}