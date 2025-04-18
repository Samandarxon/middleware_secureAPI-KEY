package middleware

import (
	"essy_travel/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSecretKeyUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UserSecretKey := ctx.GetHeader("api-key")

		if UserSecretKey != config.SecretUser {
			fmt.Println(UserSecretKey)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error_message": "Password not found"})
			return

		}
		ctx.Next()
	}
}

func CheckSecretKeyAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		AdminSecretKey := ctx.GetHeader("api-key")
		fmt.Println(AdminSecretKey)
		if AdminSecretKey != config.SecretAdmin {

			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error_message": "Password not found"})
			return

		}
		fmt.Println("AdminSecretKey")
		ctx.Next()
	}
}
