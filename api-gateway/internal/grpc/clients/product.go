package clients

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	productpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/product"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client productpb.ProductServiceClient
}

func NewProductClient(conn grpc.ClientConnInterface) *ProductClient {
	return &ProductClient{
		client: productpb.NewProductServiceClient(conn),
	}
}

func (cl *ProductClient) Create(ctx context.Context, req models.CreateProductRequest) (*productpb.Product, error) {
	resp, err := cl.client.CreateProduct(ctx, &productpb.CreateProductRequest{
		CategoryNumber:  int32(req.CategoryNumber),
		Name:            req.Name,
		Producer:        req.Producer,
		Characteristics: req.Characteristics,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetProduct(), nil
}

func (cl *ProductClient) GetByID(ctx context.Context, id int) (*productpb.Product, error) {
	resp, err := cl.client.GetProduct(ctx, &productpb.GetProductRequest{
		Id: int32(id),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetProduct(), nil
}

func (cl *ProductClient) GetAll(ctx context.Context) ([]*productpb.Product, error) {
	resp, err := cl.client.GetAllProducts(ctx, &productpb.GetAllProductsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetProducts(), nil
}

func (cl *ProductClient) GetByCategory(ctx context.Context, categoryNumber int) ([]*productpb.Product, error) {
	resp, err := cl.client.GetProductsByCategory(ctx, &productpb.GetProductsByCategoryRequest{
		CategoryNumber: int32(categoryNumber),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetProducts(), nil
}

func (cl *ProductClient) SearchByName(ctx context.Context, name string) ([]*productpb.Product, error) {
	resp, err := cl.client.SearchProductsByName(ctx, &productpb.SearchProductsByNameRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetProducts(), nil
}

func (cl *ProductClient) Update(ctx context.Context, id int, req models.UpdateProductRequest) (*productpb.Product, error) {
	resp, err := cl.client.UpdateProduct(ctx, &productpb.UpdateProductRequest{
		Id:              int32(id),
		CategoryNumber:  int32(req.CategoryNumber),
		Name:            req.Name,
		Producer:        req.Producer,
		Characteristics: req.Characteristics,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetProduct(), nil
}

func (cl *ProductClient) Delete(ctx context.Context, id int) error {
	_, err := cl.client.DeleteProduct(ctx, &productpb.DeleteProductRequest{
		Id: int32(id),
	})

	return err
}
