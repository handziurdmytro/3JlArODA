package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/routes"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080") // Or from env
}
