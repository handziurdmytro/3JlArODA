package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	categorypb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/category"
)

type CategoryClient interface {
	Create(ctx context.Context, req models.CreateCategoryRequest) (*categorypb.Category, error)
	Update(ctx context.Context, number int, req models.UpdateCategoryRequest) (*categorypb.Category, error)
	Delete(ctx context.Context, number int) error
	GetByNumber(ctx context.Context, number int) (*categorypb.Category, error)
	GetAll(ctx context.Context) ([]*categorypb.Category, error)
}

type CategoryHandler struct {
	categoryClient CategoryClient
}

func NewCategoryHandler(categoryClient CategoryClient) *CategoryHandler {
	return &CategoryHandler{categoryClient: categoryClient}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	category, err := h.categoryClient.Create(c.Request.Context(), req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetByNumber(c *gin.Context) {
	number, ok := parseIntParam(c, "number")
	if !ok {
		return
	}

	category, err := h.categoryClient.GetByNumber(c.Request.Context(), number)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.categoryClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	number, ok := parseIntParam(c, "number")
	if !ok {
		return
	}

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	category, err := h.categoryClient.Update(c.Request.Context(), number, req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	number, ok := parseIntParam(c, "number")
	if !ok {
		return
	}

	if err := h.categoryClient.Delete(c.Request.Context(), number); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
