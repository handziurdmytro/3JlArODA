package clients

import (
	"context"
	"time"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	salepb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/sale"
	"google.golang.org/grpc"
)

type SaleClient struct {
	cl salepb.SaleServiceClient
}

func NewSaleClient(conn grpc.ClientConnInterface) *SaleClient {
	return &SaleClient{
		cl: salepb.NewSaleServiceClient(conn),
	}
}

func (cl *SaleClient) Create(ctx context.Context, req models.CreateSaleRequest) error {
	_, err := cl.cl.CreateSale(ctx, &salepb.CreateSaleRequest{
		Upc:           req.UPC,
		CheckNumber:   req.CheckNumber,
		ProductNumber: int32(req.ProductNumber),
		SellingPrice:  req.SellingPrice,
	})

	return err
}

func (cl *SaleClient) GetProductSoldQuantity(ctx context.Context, productID int, from, to time.Time) (int64, error) {
	resp, err := cl.cl.GetProductSoldQuantity(ctx, &salepb.GetProductSoldQuantityRequest{
		ProductId: int32(productID),
		From:      from.Format(timeFormat),
		To:        to.Format(timeFormat),
	})
	if err != nil {
		return 0, err
	}

	return resp.GetTotalQuantity(), nil
}
