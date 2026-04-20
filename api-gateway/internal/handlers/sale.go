package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

type SaleClient interface {
	Create(ctx context.Context, req models.CreateSaleRequest) error
}

type SaleHandler struct {
	saleClient SaleClient
}

func NewSaleHandler(saleClient SaleClient) *SaleHandler {
	return &SaleHandler{saleClient: saleClient}
}

func (h *SaleHandler) Create(c *gin.Context) {
	var req models.CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	if err := h.saleClient.Create(c.Request.Context(), req); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *SaleHandler) Get(c *gin.Context) {
	upc := c.Query("upc")
	checkNumber := c.Query("check_number")

	if upc == "" || checkNumber == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC and check_number are required"})
		return
	}

	c.JSON(http.StatusNotImplemented, common.ErrorResponse{Error: "sale item lookup is not supported by business-service grpc API yet"})
}

func (h *SaleHandler) List(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, common.ErrorResponse{Error: "sale listing is not supported by business-service grpc API yet"})
}
