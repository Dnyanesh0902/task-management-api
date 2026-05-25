package middleware

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		latency := endTime.Sub(startTime)
		log.Printf(
			"%s | %s | %s | %v",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			latency,
		)
	}
}