package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Authorization Header Missing",
			})
			c.Abort()
			return 
		}
		tokenString := strings.Split(authHeader," ")

		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Invalid token format",
			})
			c.Abort()
			return 
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")),nil 
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Invalid Token",
			})
			c.Abort()
			return 
		}
		c.Next()
	}
}