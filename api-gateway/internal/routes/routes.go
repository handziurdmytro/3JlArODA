package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware()) // Example

	// Route to auth-service
	api.POST("/login", proxyToAuth)
	api.POST("/register", proxyToAuth)

	// Route to business-service
	api.GET("/employees", proxyToBusiness)
	// Add more routes
}

func proxyToAuth(c *gin.Context) {
	// Implement proxy to auth-service
	c.JSON(200, gin.H{"message": "Proxy to auth"})
}

func proxyToBusiness(c *gin.Context) {
	// Implement proxy to business-service
	c.JSON(200, gin.H{"message": "Proxy to business"})
}
