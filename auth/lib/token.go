package lib

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Parse token string and return payload
func ParseToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

// Extract token from request Authorization header
func ExtractTokenFromRequest(c *gin.Context) string {
	if c.Request.Header.Get("Authorization") == "" {
		return ""
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
