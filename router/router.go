package router

import (
	"exchangedapp_backend/controllers"
	"exchangedapp_backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com" //动态决定允许什么域名访问，可忽略掉AllowOrigins
		},
		MaxAge: 12 * time.Hour, //该时间内只需要发送一次OPTIONS prefile请求
	}))

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
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticlesByID)

		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("/articles/:id/like", controllers.GetArticleLikes)
	}
	return r
}
