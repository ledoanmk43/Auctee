package middleware

import (
	"chilindo/pkg/token"
	"chilindo/src/user-service/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type IMiddleWare interface {
	MiddleWare() gin.HandlerFunc
}

type SMiddleWare struct {
	tokenController *token.Claims
}

func (s *SMiddleWare) IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "Request doest not contain token",
			})
			log.Println("MiddleWare: Error to get token in")
			c.Abort()
			return
		}
		tokenResult := strings.TrimPrefix(tokenString, "Bearer ")
		claim, err := token.ExtractToken(tokenResult)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": err.Error(),
			})
			log.Println("Error:", err.Error())
			c.Abort()
			return
		}
		c.Set(config.UserId, claim.UserId)
		if claim.ExpiresAt < time.Now().Local().Unix() {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "token is expired",
			})
			log.Println("MiddleWare: Error token is expired")
			c.Abort()
			return
		}
		c.Next()
	}
}
