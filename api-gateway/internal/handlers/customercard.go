package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

func CreateCustomerCard(c *gin.Context) {
	var req models.CreateCustomerCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Pass to the business service via gRPC
	c.JSON(http.StatusCreated, req)
}

func GetCustomerCardByNumber(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Card number is required"})
		return
	}

	// TODO: Fetch from the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"card_number": number, "status": "found"})
}

func ListCustomerCards(c *gin.Context) {
	// TODO: Fetch collection from the business service via gRPC
	c.JSON(http.StatusOK, []string{})
}

func UpdateCustomerCard(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Card number is required"})
		return
	}

	var req models.UpdateCustomerCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	// TODO: Send update to the business service via gRPC
	c.JSON(http.StatusOK, gin.H{"card_number": number, "status": "updated"})
}

func DeleteCustomerCard(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Card number is required"})
		return
	}

	// TODO: Send delete to the business service via gRPC
	c.Status(http.StatusNoContent)
}
