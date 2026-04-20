package clients

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	storeproductpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/storeproduct"
	"google.golang.org/grpc"
)

type StoreProductClient struct {
	client storeproductpb.StoreProductServiceClient
}

func NewStoreProductClient(conn grpc.ClientConnInterface) *StoreProductClient {
	return &StoreProductClient{
		client: storeproductpb.NewStoreProductServiceClient(conn),
	}
}

func (cl *StoreProductClient) Create(ctx context.Context, req models.CreateStoreProductRequest) (*storeproductpb.StoreProduct, error) {
	resp, err := cl.client.CreateStoreProduct(ctx, &storeproductpb.CreateStoreProductRequest{
		Upc:                req.UPC,
		UpcProm:            req.UPCProm,
		ProductId:          int32(req.ProductID),
		SellingPrice:       req.SellingPrice,
		ProductsNumber:     int32(req.ProductsNumber),
		PromotionalProduct: req.PromotionalProduct,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProduct(), nil
}

func (cl *StoreProductClient) Update(ctx context.Context, upc string, req models.UpdateStoreProductRequest) (*storeproductpb.StoreProduct, error) {
	resp, err := cl.client.UpdateStoreProduct(ctx, &storeproductpb.UpdateStoreProductRequest{
		Upc:                upc,
		UpcProm:            req.UPCProm,
		ProductId:          int32(req.ProductID),
		SellingPrice:       req.SellingPrice,
		ProductsNumber:     int32(req.ProductsNumber),
		PromotionalProduct: req.PromotionalProduct,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProduct(), nil
}

func (cl *StoreProductClient) Delete(ctx context.Context, upc string) error {
	_, err := cl.client.DeleteStoreProduct(ctx, &storeproductpb.DeleteStoreProductRequest{Upc: upc})
	return err
}

func (cl *StoreProductClient) GetByUPC(ctx context.Context, upc string) (*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetStoreProduct(ctx, &storeproductpb.GetStoreProductRequest{Upc: upc})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProduct(), nil
}

func (cl *StoreProductClient) GetAllSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetAllStoreProductsSortedByQuantity(ctx, &storeproductpb.GetAllStoreProductsSortedByQuantityRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetAllSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetAllStoreProductsSortedByName(ctx, &storeproductpb.GetAllStoreProductsSortedByNameRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetPromoSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetPromoStoreProductsSortedByQuantity(ctx, &storeproductpb.GetPromoStoreProductsSortedByQuantityRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetPromoSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetPromoStoreProductsSortedByName(ctx, &storeproductpb.GetPromoStoreProductsSortedByNameRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetNonPromoSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetNonPromoStoreProductsSortedByQuantity(ctx, &storeproductpb.GetNonPromoStoreProductsSortedByQuantityRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetNonPromoSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetNonPromoStoreProductsSortedByName(ctx, &storeproductpb.GetNonPromoStoreProductsSortedByNameRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}

func (cl *StoreProductClient) GetByCategorySortedByName(ctx context.Context, categoryNumber int) ([]*storeproductpb.DetailedStoreProduct, error) {
	resp, err := cl.client.GetStoreProductsByCategorySortedByName(ctx, &storeproductpb.GetStoreProductsByCategorySortedByNameRequest{
		CategoryNumber: int32(categoryNumber),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetStoreProducts(), nil
}
