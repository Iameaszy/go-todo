package middlewares

import (
	"strings"
	"todo/modules/Auth/AuthService"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		splittedToken := strings.Split(tokenString, " ")
		if len(splittedToken) < 2 || splittedToken[1] == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err := AuthService.ValidateToken(splittedToken[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
