package sale

import (
	"context"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	salepb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/sale"
)

type GRPCHandler struct {
	service *Service
	salepb.UnimplementedSaleServiceServer
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateSale(ctx context.Context, req *salepb.CreateSaleRequest) (*salepb.CreateSaleResponse, error) {
	err := h.service.Create(ctx, CreateRequest{
		UPC:           req.GetUpc(),
		CheckNumber:   req.GetCheckNumber(),
		ProductNumber: int(req.GetProductNumber()),
		SellingPrice:  req.GetSellingPrice(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &salepb.CreateSaleResponse{Success: true}, nil
}

func (h *GRPCHandler) GetProductSoldQuantity(ctx context.Context, req *salepb.GetProductSoldQuantityRequest) (*salepb.GetProductSoldQuantityResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	quantity, err := h.service.GetProductSoldQuantity(ctx, int(req.GetProductId()), from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &salepb.GetProductSoldQuantityResponse{TotalQuantity: quantity}, nil
}

func parseProtoPeriod(fromValue, toValue string) (from time.Time, to time.Time, err error) {
	from, err = common.ParseProtoTime(fromValue)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	to, err = common.ParseProtoTime(toValue)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return from, to, nil
}
