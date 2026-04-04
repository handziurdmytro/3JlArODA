package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

func CreateStoreProduct(c *gin.Context) {
	var req models.CreateStoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Pass to the business service via gRPC
	c.JSON(http.StatusCreated, req)
}

func GetStoreProductByUPC(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	// TODO: Fetch from the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"upc": upc, "status": "found"})
}

func ListStoreProducts(c *gin.Context) {
	// TODO: Fetch collection from the business service via gRPC
	c.JSON(http.StatusOK, []string{})
}

func UpdateStoreProduct(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	var req models.UpdateStoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Send update to the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"upc": upc, "status": "updated"})
}

func DeleteStoreProduct(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	// TODO: Send delete to the business service via gRPC
	c.Status(http.StatusNoContent)
}
