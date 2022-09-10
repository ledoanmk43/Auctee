package utils

import (
	"backend/pkg/token"
	"errors"
	"github.com/gin-contrib/sessions"
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
			"message": "Unauthorized",
		})
		ctx.Abort()
		return 0
	}

	id := claims.UserId
	return id
}

func GetTokenFromCookie(ctx *gin.Context, cookieName string) (string, error) {
	newSession := sessions.DefaultMany(ctx, cookieName)
	tokenFromCookie := newSession.Get(cookieName)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return "", errors.New("no cookie")
	}
	return tokenFromCookie.(string), nil
}
