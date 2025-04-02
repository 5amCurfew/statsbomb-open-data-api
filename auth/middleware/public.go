package middleware

import (
	"github.com/gin-gonic/gin"
)

func PublicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
