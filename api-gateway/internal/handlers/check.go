package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	checkpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/check"
)

type CheckClient interface {
	Create(ctx context.Context, req models.CreateCheckRequest) error
	Delete(ctx context.Context, number string) error
	GetAll(ctx context.Context) ([]*checkpb.Check, error)
	GetOfDayByCashier(ctx context.Context, employeeID string, date time.Time) ([]*checkpb.Check, error)
	GetOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]*checkpb.Check, error)
	GetFullData(ctx context.Context, number string) ([]*checkpb.FullCheckItem, error)
	GetDetailsOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]*checkpb.CheckDetail, error)
	GetDetailsOfPeriod(ctx context.Context, from, to time.Time) ([]*checkpb.CheckDetail, error)
	GetSumOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) (float64, error)
	GetSumOfPeriod(ctx context.Context, from, to time.Time) (float64, error)
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
	cashierID := c.Query("cashier_id")
	date := c.Query("date")
	if cashierID != "" && date != "" {
		day, ok := parseTimeQuery(c, "date")
		if !ok {
			return
		}

		checks, err := h.checkClient.GetOfDayByCashier(c.Request.Context(), cashierID, day)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, checks)
		return
	}

	fromRaw := c.Query("from")
	toRaw := c.Query("to")
	if cashierID != "" && fromRaw != "" && toRaw != "" {
		from, to, ok := parsePeriod(c)
		if !ok {
			return
		}

		checks, err := h.checkClient.GetOfPeriodByCashier(c.Request.Context(), cashierID, from, to)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, checks)
		return
	}

	checks, err := h.checkClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, checks)
}

func (h *CheckHandler) GetDetailsReport(c *gin.Context) {
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	cashierID := c.Query("cashier_id")
	if cashierID != "" {
		details, err := h.checkClient.GetDetailsOfPeriodByCashier(c.Request.Context(), cashierID, from, to)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, details)
		return
	}

	details, err := h.checkClient.GetDetailsOfPeriod(c.Request.Context(), from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, details)
}

func (h *CheckHandler) GetTotalReport(c *gin.Context) {
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	cashierID := c.Query("cashier_id")
	if cashierID != "" {
		total, err := h.checkClient.GetSumOfPeriodByCashier(c.Request.Context(), cashierID, from, to)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"total_sum": total})
		return
	}

	total, err := h.checkClient.GetSumOfPeriod(c.Request.Context(), from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_sum": total})
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

func parsePeriod(c *gin.Context) (time.Time, time.Time, bool) {
	from, ok := parseTimeQuery(c, "from")
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	to, ok := parseTimeQuery(c, "to")
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	return from, to, true
}
