package sale

import salepb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/sale"

type GRPCHandler struct {
	service *Service
	salepb.UnimplementedSaleServiceServer
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}
