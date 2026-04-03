package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check token, validate with auth-service
		// For now, pass
		c.Next()
	}
}
