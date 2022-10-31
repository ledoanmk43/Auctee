package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request from %s: %s", c.ClientIP(), c.Request.URL.Path)
		c.Next()
	}
}
