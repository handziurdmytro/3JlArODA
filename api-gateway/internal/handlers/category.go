package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Pass to the business service via gRPC
	c.JSON(http.StatusCreated, req)
}

func GetCategoryByNumber(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Number is required"})
		return
	}

	// TODO: Fetch from the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"number": number, "status": "found"})
}

func ListCategories(c *gin.Context) {
	// TODO: Fetch collection from the business service via gRPC
	c.JSON(http.StatusOK, []string{})
}

func UpdateCategory(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Number is required"})
		return
	}

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Send update to the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"number": number, "status": "updated"})
}

func DeleteCategory(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Number is required"})
		return
	}

	// TODO: Send delete to the business service via gRPC
	c.Status(http.StatusNoContent)
}
