package middleware

import (
	"github.com/gin-gonic/gin"
)

func Public() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
