package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "../frontend-web/dist/assets")
	router.Static("/public", "../frontend-web/public")
	router.POST("/api/register", register)
	router.NoRoute(root)

	router.Run(":7373")
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
