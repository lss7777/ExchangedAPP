package middlewares

import (
	"exchangedapp_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "MissingAuthorization"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invaild Token"})
		}
		ctx.Set("username", username)

	}
}
