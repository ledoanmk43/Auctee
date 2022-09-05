package utils

import (
	"chilindo/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Router() *gin.Engine {
	router := gin.Default()
	return router
}

func GetIdFromToken(ctx *gin.Context) uint {
	rawToken := ctx.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(rawToken, "Bearer ")
	claims, errExtract := token.ExtractToken(tokenString)

	if errExtract != nil || len(tokenString) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Unauthorized",
		})
		ctx.Abort()
		return 0
	}

	adminId := claims.UserId
	return adminId
}
