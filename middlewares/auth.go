package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest-api.com/m/v2/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized1"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		println(err.Error())
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized2"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
