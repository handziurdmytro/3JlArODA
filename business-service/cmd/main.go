package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/database/postgres"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	pool, err := postgres.GetNewPool(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	router := gin.Default()

	router.GET("/employees", handlers.GetEmployees(pool))
	router.POST("/employees", handlers.CreateEmployee(pool))

	router.Run(":8082")
}
