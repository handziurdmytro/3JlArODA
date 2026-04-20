package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
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
}

type EmployeeHandler struct {
	employeeClient EmployeeClient
}

func NewEmployeeHandler(employeeClient EmployeeClient) *EmployeeHandler {
	return &EmployeeHandler{employeeClient: employeeClient}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req models.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	employee, err := h.employeeClient.Create(c.Request.Context(), req)
	if err != nil {
		respondGRPCError(c, err)
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
