package storeproduct

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	storeproductpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/storeproduct"
)

type GRPCHandler struct {
	storeproductpb.UnimplementedStoreProductServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateStoreProduct(ctx context.Context, req *storeproductpb.CreateStoreProductRequest) (*storeproductpb.StoreProductResponse, error) {
	product, err := h.service.Create(ctx, CreateRequest{
		UPC:                req.GetUpc(),
		UPCProm:            req.UpcProm,
		ProductID:          int(req.GetProductId()),
		SellingPrice:       req.GetSellingPrice(),
		ProductsNumber:     int(req.GetProductsNumber()),
		PromotionalProduct: req.GetPromotionalProduct(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.StoreProductResponse{StoreProduct: toProtoStoreProduct(product)}, nil
}

func (h *GRPCHandler) UpdateStoreProduct(ctx context.Context, req *storeproductpb.UpdateStoreProductRequest) (*storeproductpb.StoreProductResponse, error) {
	product, err := h.service.Update(ctx, StoreProduct{
		UPC:                req.GetUpc(),
		UPCProm:            req.UpcProm,
		ProductID:          int(req.GetProductId()),
		SellingPrice:       req.GetSellingPrice(),
		ProductsNumber:     int(req.GetProductsNumber()),
		PromotionalProduct: req.GetPromotionalProduct(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.StoreProductResponse{StoreProduct: toProtoStoreProduct(product)}, nil
}

func (h *GRPCHandler) DeleteStoreProduct(ctx context.Context, req *storeproductpb.DeleteStoreProductRequest) (*storeproductpb.DeleteStoreProductResponse, error) {
	if err := h.service.Delete(ctx, req.GetUpc()); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.DeleteStoreProductResponse{Success: true}, nil
}

func (h *GRPCHandler) GetStoreProduct(ctx context.Context, req *storeproductpb.GetStoreProductRequest) (*storeproductpb.DetailedStoreProductResponse, error) {
	product, err := h.service.GetByUPC(ctx, req.GetUpc())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.DetailedStoreProductResponse{StoreProduct: toProtoDetailedStoreProduct(product)}, nil
}

func (h *GRPCHandler) GetAllStoreProductsSortedByQuantity(ctx context.Context, _ *storeproductpb.GetAllStoreProductsSortedByQuantityRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetAllSortedByQuantity(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetAllStoreProductsSortedByName(ctx context.Context, _ *storeproductpb.GetAllStoreProductsSortedByNameRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetAllSortedByName(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetPromoStoreProductsSortedByQuantity(ctx context.Context, _ *storeproductpb.GetPromoStoreProductsSortedByQuantityRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetPromoSortedByQuantity(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetPromoStoreProductsSortedByName(ctx context.Context, _ *storeproductpb.GetPromoStoreProductsSortedByNameRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetPromoSortedByName(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetNonPromoStoreProductsSortedByQuantity(ctx context.Context, _ *storeproductpb.GetNonPromoStoreProductsSortedByQuantityRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetNonPromoSortedByQuantity(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetNonPromoStoreProductsSortedByName(ctx context.Context, _ *storeproductpb.GetNonPromoStoreProductsSortedByNameRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetNonPromoSortedByName(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func (h *GRPCHandler) GetStoreProductsByCategorySortedByName(ctx context.Context, req *storeproductpb.GetStoreProductsByCategorySortedByNameRequest) (*storeproductpb.GetDetailedStoreProductsResponse, error) {
	products, err := h.service.GetByCategorySortedByName(ctx, int(req.GetCategoryNumber()))
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &storeproductpb.GetDetailedStoreProductsResponse{StoreProducts: toProtoDetailedStoreProducts(products)}, nil
}

func toProtoStoreProduct(product *StoreProduct) *storeproductpb.StoreProduct {
	if product == nil {
		return nil
	}

	return &storeproductpb.StoreProduct{
		Upc:                product.UPC,
		UpcProm:            product.UPCProm,
		ProductId:          int32(product.ProductID),
		SellingPrice:       product.SellingPrice,
		ProductsNumber:     int32(product.ProductsNumber),
		PromotionalProduct: product.PromotionalProduct,
	}
}

func toProtoDetailedStoreProducts(products []DetailedStoreProduct) []*storeproductpb.DetailedStoreProduct {
	result := make([]*storeproductpb.DetailedStoreProduct, 0, len(products))
	for _, product := range products {
		product := product
		result = append(result, toProtoDetailedStoreProduct(&product))
	}

	return result
}

func toProtoDetailedStoreProduct(product *DetailedStoreProduct) *storeproductpb.DetailedStoreProduct {
	if product == nil {
		return nil
	}

	return &storeproductpb.DetailedStoreProduct{
		Upc:                product.UPC,
		UpcProm:            product.UPCProm,
		SellingPrice:       product.SellingPrice,
		ProductsNumber:     int32(product.ProductsNumber),
		PromotionalProduct: product.PromotionalProduct,
		ProductId:          int32(product.ProductID),
		ProductName:        product.ProductName,
		Producer:           product.Producer,
		Characteristics:    product.Characteristics,
		CategoryNumber:     int32(product.CategoryNumber),
		CategoryName:       product.CategoryName,
	}
}
