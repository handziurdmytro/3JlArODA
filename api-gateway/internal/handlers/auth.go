package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	"github.com/handziurdmytro/3JlArODA/api-gateway/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthClient interface {
	Register(ctx context.Context, username, password string) (*pb.RegisterResponse, error)
	Login(ctx context.Context, username, password string) (*pb.LoginResponse, error)
}

type AuthHandler struct {
	authClient AuthClient
}

func NewAuthHandler(authClient AuthClient) *AuthHandler {
	return &AuthHandler{authClient: authClient}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	resp, err := h.authClient.Register(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		switch status.Code(err) {
		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, common.ErrorResponse{Error: "User already exists"})
		default:
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Failed to register"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": resp.Token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	resp, err := h.authClient.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		switch status.Code(err) {
		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "Invalid credentials"})
		case codes.NotFound:
			c.JSON(http.StatusNotFound, common.ErrorResponse{Error: "User not found"})
		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, common.ErrorResponse{Error: "User already exists"})
		default:
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Failed to login"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}
