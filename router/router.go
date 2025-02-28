package router

import (
	"exchangedapp_backend/controllers"
	"exchangedapp_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("login", controllers.Login)

		auth.POST("register", controllers.Register)

	}

	api := r.Group("/api")
	api.GET("exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("exchangeRates", controllers.CreateExchangeRate)
	}
	return r
}
