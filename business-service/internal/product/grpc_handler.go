package product

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	productpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/product"
)

type GRPCHandler struct {
	productpb.UnimplementedProductServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.ProductResponse, error) {
	product, err := h.service.Create(ctx, CreateRequest{
		CategoryNumber:  int(req.GetCategoryNumber()),
		Name:            req.GetName(),
		Producer:        req.Producer,
		Characteristics: req.GetCharacteristics(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.ProductResponse{Product: toProtoProduct(product)}, nil
}

func (h *GRPCHandler) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.ProductResponse, error) {
	product, err := h.service.Update(ctx, Product{
		ID:              int(req.GetId()),
		CategoryNumber:  int(req.GetCategoryNumber()),
		Name:            req.GetName(),
		Producer:        req.Producer,
		Characteristics: req.GetCharacteristics(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.ProductResponse{Product: toProtoProduct(product)}, nil
}

func (h *GRPCHandler) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	if err := h.service.Delete(ctx, int(req.GetId())); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.DeleteProductResponse{Success: true}, nil
}

func (h *GRPCHandler) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.ProductResponse, error) {
	product, err := h.service.GetByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.ProductResponse{Product: toProtoProduct(product)}, nil
}

func (h *GRPCHandler) GetAllProducts(ctx context.Context, _ *productpb.GetAllProductsRequest) (*productpb.GetProductsResponse, error) {
	products, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.GetProductsResponse{Products: toProtoProducts(products)}, nil
}

func (h *GRPCHandler) GetProductsByCategory(ctx context.Context, req *productpb.GetProductsByCategoryRequest) (*productpb.GetProductsResponse, error) {
	products, err := h.service.GetByCategory(ctx, int(req.GetCategoryNumber()))
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.GetProductsResponse{Products: toProtoProducts(products)}, nil
}

func (h *GRPCHandler) SearchProductsByName(ctx context.Context, req *productpb.SearchProductsByNameRequest) (*productpb.GetProductsResponse, error) {
	products, err := h.service.SearchByName(ctx, req.GetName())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &productpb.GetProductsResponse{Products: toProtoProducts(products)}, nil
}

func toProtoProducts(products []Product) []*productpb.Product {
	result := make([]*productpb.Product, 0, len(products))
	for _, product := range products {
		product := product
		result = append(result, toProtoProduct(&product))
	}

	return result
}

func toProtoProduct(product *Product) *productpb.Product {
	if product == nil {
		return nil
	}

	return &productpb.Product{
		Id:              int32(product.ID),
		CategoryNumber:  int32(product.CategoryNumber),
		Name:            product.Name,
		Producer:        product.Producer,
		Characteristics: product.Characteristics,
	}
}
