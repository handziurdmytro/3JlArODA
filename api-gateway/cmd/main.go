package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/routes"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
