package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	authpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb"
	employeepb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/employee"
)

type EmployeeClient interface {
	Create(ctx context.Context, req models.CreateEmployeeRequest) (*employeepb.Employee, error)
	Update(ctx context.Context, id string, req models.UpdateEmployeeRequest) (*employeepb.Employee, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*employeepb.Employee, error)
	GetAll(ctx context.Context) ([]*employeepb.Employee, error)
	GetByRole(ctx context.Context, role string) ([]*employeepb.Employee, error)
	GetContactsBySurname(ctx context.Context, surname string) ([]*employeepb.EmployeeContact, error)
	GetContactsByFullName(ctx context.Context, surname, name string, patronymic *string) ([]*employeepb.EmployeeContact, error)
	GetCashierPerformance(ctx context.Context, from, to time.Time, minRevenue float64) ([]*employeepb.CashierPerformance, error)
	GetBestCashiersByPromo(ctx context.Context) ([]*employeepb.BestCashierByPromo, error)
}

type EmployeeHandler struct {
	authClient     AuthClient
	employeeClient EmployeeClient
}

func NewEmployeeHandler(employeeClient EmployeeClient, authClient AuthClient) *EmployeeHandler {
	return &EmployeeHandler{
		employeeClient: employeeClient,
		authClient:     authClient,
	}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req models.CreateFullEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	_, err := h.authClient.Register(c.Request.Context(), &authpb.RegisterRequest{
		Id:       req.EmployeeData.ID,
		Username: req.AuthData.Username,
		Password: req.AuthData.Password,
		Role:     req.EmployeeData.Role,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create security account: " + err.Error()})
		return
	}

	businessReq := models.CreateEmployeeRequest{
		ID:          req.EmployeeData.ID,
		Surname:     req.EmployeeData.Surname,
		Name:        req.EmployeeData.Name,
		Patronymic:  req.EmployeeData.Patronymic,
		Role:        req.EmployeeData.Role,
		Salary:      req.EmployeeData.Salary,
		DateOfBirth: req.EmployeeData.DateOfBirth,
		DateOfStart: req.EmployeeData.DateOfStart,
		PhoneNumber: req.EmployeeData.PhoneNumber,
		City:        req.EmployeeData.City,
		Street:      req.EmployeeData.Street,
		ZipCode:     req.EmployeeData.ZipCode,
	}

	employee, err := h.employeeClient.Create(c.Request.Context(), businessReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Account created, but business profile failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	employee, err := h.employeeClient.GetByID(c.Request.Context(), id)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) GetMe(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "missing authenticated user"})
		return
	}

	id, ok := userID.(string)
	if !ok || id == "" {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "invalid authenticated user"})
		return
	}

	employee, err := h.employeeClient.GetByID(c.Request.Context(), id)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) List(c *gin.Context) {
	role := c.Query("role")
	if role != "" {
		employees, err := h.employeeClient.GetByRole(c.Request.Context(), role)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, employees)
		return
	}

	employees, err := h.employeeClient.GetAll(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetContacts(c *gin.Context) {
	surname := c.Query("surname")
	if surname == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "surname is required"})
		return
	}

	name := c.Query("name")
	if name == "" {
		contacts, err := h.employeeClient.GetContactsBySurname(c.Request.Context(), surname)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, contacts)
		return
	}

	var patronymic *string
	if value := c.Query("patronymic"); value != "" {
		patronymic = &value
	}

	contacts, err := h.employeeClient.GetContactsByFullName(c.Request.Context(), surname, name, patronymic)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func (h *EmployeeHandler) GetCashierPerformance(c *gin.Context) {
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}
	minRevenue, ok := parseFloatQuery(c, "min_revenue")
	if !ok {
		return
	}

	performance, err := h.employeeClient.GetCashierPerformance(c.Request.Context(), from, to, minRevenue)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, performance)
}

func (h *EmployeeHandler) GetBestCashiersByPromo(c *gin.Context) {
	cashiers, err := h.employeeClient.GetBestCashiersByPromo(c.Request.Context())
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, cashiers)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	var req models.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	employee, err := h.employeeClient.Update(c.Request.Context(), id, req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID is required"})
		return
	}

	if err := h.employeeClient.Delete(c.Request.Context(), id); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
