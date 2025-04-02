package middleware

import (
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data-api/auth/lib"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Check if token is authorised for admin routes
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authorised bool = false

		tokenString := lib.ExtractTokenFromRequest(c)
		token, _ := lib.ParseToken(tokenString)

		if token.Valid {
			claims, _ := token.Claims.(jwt.MapClaims)
			authorised = claims["adm"].(bool)
		}

		if authorised {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
			c.Abort()
			return
		}
	}
}
