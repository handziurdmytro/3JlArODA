package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	customercardpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/customercard"
)

type CustomerCardClient interface {
	Create(ctx context.Context, req models.CreateCustomerCardRequest) (*customercardpb.CustomerCard, error)
	GetByNumber(ctx context.Context, cardNumber string) (*customercardpb.CustomerCard, error)
	GetAll(ctx context.Context) ([]*customercardpb.CustomerCard, error)
	GetByPercent(ctx context.Context, percent int) ([]*customercardpb.CustomerCard, error)
	SearchBySurname(ctx context.Context, surname string) ([]*customercardpb.CustomerCard, error)
	GetWhoBoughtAllProductsFromCategory(ctx context.Context, categoryNumber int, from, to time.Time) ([]*customercardpb.CustomerCard, error)
	Update(ctx context.Context, cardNumber string, req models.UpdateCustomerCardRequest) (*customercardpb.CustomerCard, error)
	Delete(ctx context.Context, cardNumber string) error
}

type CustomerCardHandler struct {
	customerCardClient CustomerCardClient
}

func NewCustomerCardHandler(customerCardClient CustomerCardClient) *CustomerCardHandler {
	return &CustomerCardHandler{customerCardClient: customerCardClient}
}

func (h *CustomerCardHandler) Create(c *gin.Context) {
	var req models.CreateCustomerCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	card, err := h.customerCardClient.Create(c.Request.Context(), req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, card)
}

func (h *CustomerCardHandler) GetByNumber(c *gin.Context) {
	number, ok := stringParam(c, "number", "Card number is required")
	if !ok {
		return
	}

	card, err := h.customerCardClient.GetByNumber(c.Request.Context(), number)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, card)
}

func (h *CustomerCardHandler) List(c *gin.Context) {
	surname := c.Query("surname")
	if surname != "" {
		cards, err := h.customerCardClient.SearchBySurname(c.Request.Context(), surname)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, cards)
		return
	}

	percent := c.Query("percent")
	if percent != "" {
		parsed, err := strconv.Atoi(percent)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid percent"})
			return
		}

		cards, err := h.customerCardClient.GetByPercent(c.Request.Context(), parsed)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, cards)
		return
	}

	cards, err := h.customerCardClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, cards)
}

func (h *CustomerCardHandler) GetWhoBoughtAllProductsFromCategory(c *gin.Context) {
	categoryNumber, ok := parseIntQuery(c, "category_number")
	if !ok {
		return
	}
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	cards, err := h.customerCardClient.GetWhoBoughtAllProductsFromCategory(c.Request.Context(), categoryNumber, from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, cards)
}

func (h *CustomerCardHandler) Update(c *gin.Context) {
	number, ok := stringParam(c, "number", "Card number is required")
	if !ok {
		return
	}

	var req models.UpdateCustomerCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	card, err := h.customerCardClient.Update(c.Request.Context(), number, req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, card)
}

func (h *CustomerCardHandler) Delete(c *gin.Context) {
	number, ok := stringParam(c, "number", "Card number is required")
	if !ok {
		return
	}

	if err := h.customerCardClient.Delete(c.Request.Context(), number); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func stringParam(c *gin.Context, name, errorMessage string) (string, bool) {
	value := c.Param(name)
	if value == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: errorMessage})
		return "", false
	}

	return value, true
}
