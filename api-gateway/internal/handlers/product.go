package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

func CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Pass to the business service via gRPC
	c.JSON(http.StatusCreated, req)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	// TODO: Fetch from the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"id": id, "status": "found"})
}

func ListProducts(c *gin.Context) {
	// TODO: Fetch collection from the business service via gRPC
	c.JSON(http.StatusOK, []string{})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Send update to the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"id": id, "status": "updated"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	// TODO: Send delete to the business service via gRPC
	c.Status(http.StatusNoContent)
}
