package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromRequest(c)
		token, _ := parseToken(tokenString)

		if token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
			c.Abort()
			return
		}
	}
}

// Parse token string and return payload
func parseToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

// Extract token from request Authorization header
func extractTokenFromRequest(c *gin.Context) string {
	if c.Request.Header.Get("Authorization") == "" {
		return ""
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
