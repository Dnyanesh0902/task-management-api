package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRole(role string) gin.HandlerFunc {
	return  func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Authorization header missing",
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
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil ||!token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"invalid token",
			})
			c.Abort()
			return 
		}

		claims := token.Claims.(jwt.MapClaims)

		userRole := claims["role"].(string)

		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{
				"message":"Access denied",

			})
			c.Abort()
			return 
		}
		c.Next()
	}
}