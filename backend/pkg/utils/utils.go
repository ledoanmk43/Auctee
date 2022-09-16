package utils

import (
	"backend/pkg/token"
	"backend/src/auction-service/config"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return "", errors.New("no cookie")
	}
	return tokenFromCookie.(string), nil
}

func GetMoment() (time.Time, error) {
	now, err := time.Parse(config.DATEFORMAT, time.Now().Format(config.DATEFORMAT))
	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

func StringToTime(timeString string) (time.Time, error) {
	result, err := time.Parse(config.DATEFORMAT, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
