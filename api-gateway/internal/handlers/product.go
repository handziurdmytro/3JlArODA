package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	productpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/product"
)

type ProductClient interface {
	Create(ctx context.Context, req models.CreateProductRequest) (*productpb.Product, error)
	GetByID(ctx context.Context, id int) (*productpb.Product, error)
	GetAll(ctx context.Context) ([]*productpb.Product, error)
	GetByCategory(ctx context.Context, categoryNumber int) ([]*productpb.Product, error)
	SearchByName(ctx context.Context, name string) ([]*productpb.Product, error)
	GetSalesStatsByCategoryAndPeriod(ctx context.Context, categoryNumber int, from, to time.Time) ([]*productpb.ProductSalesStats, error)
	Update(ctx context.Context, id int, req models.UpdateProductRequest) (*productpb.Product, error)
	Delete(ctx context.Context, id int) error
}

type ProductHandler struct {
	productClient ProductClient
}

func NewProductHandler(productClient ProductClient) *ProductHandler {
	return &ProductHandler{productClient: productClient}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	product, err := h.productClient.Create(c.Request.Context(), req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	product, err := h.productClient.GetByID(c.Request.Context(), id)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) List(c *gin.Context) {
	name := c.Query("name")
	if name != "" {
		products, err := h.productClient.SearchByName(c.Request.Context(), name)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, products)
		return
	}

	categoryNumber := c.Query("category_number")
	if categoryNumber != "" {
		parsed, err := strconv.Atoi(categoryNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid category_number"})
			return
		}

		products, err := h.productClient.GetByCategory(c.Request.Context(), parsed)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, products)
		return
	}

	products, err := h.productClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetSalesStatsByCategoryAndPeriod(c *gin.Context) {
	categoryNumber, ok := parseIntQuery(c, "category_number")
	if !ok {
		return
	}
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	stats, err := h.productClient.GetSalesStatsByCategoryAndPeriod(c.Request.Context(), categoryNumber, from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	product, err := h.productClient.Update(c.Request.Context(), id, req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, ok := parseIntParam(c, "id")
	if !ok {
		return
	}

	if err := h.productClient.Delete(c.Request.Context(), id); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func parseIntParam(c *gin.Context, name string) (int, bool) {
	value := c.Param(name)
	if value == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: name + " is required"})
		return 0, false
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid " + name})
		return 0, false
	}

	return parsed, true
}
