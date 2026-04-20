package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
)

type SaleClient interface {
	Create(ctx context.Context, req models.CreateSaleRequest) error
	GetProductSoldQuantity(ctx context.Context, productID int, from, to time.Time) (int64, error)
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

	if number := c.Param("number"); number != "" {
		req.CheckNumber = number
	}
	if req.CheckNumber == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "check_number is required"})
		return
	}

	if err := h.saleClient.Create(c.Request.Context(), req); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *SaleHandler) GetProductSoldQuantity(c *gin.Context) {
	productID, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	quantity, err := h.saleClient.GetProductSoldQuantity(c.Request.Context(), productID, from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID, "total_quantity": quantity})
}
