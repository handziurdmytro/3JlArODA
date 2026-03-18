package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/go-service/internal/database/postgres"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	pool, err := postgres.GetNewPool(context.Background(), dsn)
	_ = pool
	_ = err

	router := gin.Default()

	router.Static("/assets", "../frontend-web/dist/assets")
	router.Static("/public", "../frontend-web/public")
	router.POST("/api/register", register)
	router.NoRoute(root)

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}

func root(c *gin.Context) {
	c.File("../frontend-web/dist/index.html")
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	var req RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	resp := RegisterResponse{
		Email:    req.Email,
		Password: req.Password,
	}

	fmt.Println(resp.Email)
	fmt.Println(resp.Password)

	c.JSON(http.StatusOK, resp)
}
