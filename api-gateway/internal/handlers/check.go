package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	checkpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/check"
)

type CheckClient interface {
	Create(ctx context.Context, req models.CreateCheckRequest) error
	Delete(ctx context.Context, number string) error
	GetAll(ctx context.Context) ([]*checkpb.Check, error)
	GetFullData(ctx context.Context, number string) ([]*checkpb.FullCheckItem, error)
}

type CheckHandler struct {
	checkClient CheckClient
}

func NewCheckHandler(checkClient CheckClient) *CheckHandler {
	return &CheckHandler{checkClient: checkClient}
}

func (h *CheckHandler) Create(c *gin.Context) {
	var req models.CreateCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	if err := h.checkClient.Create(c.Request.Context(), req); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *CheckHandler) GetByNumber(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Number is required"})
		return
	}

	items, err := h.checkClient.GetFullData(c.Request.Context(), number)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *CheckHandler) List(c *gin.Context) {
	checks, err := h.checkClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, checks)
}

func (h *CheckHandler) Delete(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Number is required"})
		return
	}

	if err := h.checkClient.Delete(c.Request.Context(), number); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
