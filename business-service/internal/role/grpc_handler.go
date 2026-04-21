package role

import (
	"context"

	rolepb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/role"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	rolepb.UnimplementedRoleServiceServer
	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) GetEmployeeRole(ctx context.Context, req *rolepb.GetEmployeeRoleRequest) (*rolepb.GetEmployeeRoleResponse, error) {
	role, err := h.service.GetEmployeeRole(ctx, req.EmployeeId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	return &rolepb.GetEmployeeRoleResponse{
		Role: role,
	}, nil
}
