package clients

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	categorypb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/category"
	"google.golang.org/grpc"
)

type CategoryClient struct {
	client categorypb.CategoryServiceClient
}

func NewCategoryClient(conn grpc.ClientConnInterface) *CategoryClient {
	return &CategoryClient{
		client: categorypb.NewCategoryServiceClient(conn),
	}
}

func (cl *CategoryClient) Create(ctx context.Context, req models.CreateCategoryRequest) (*categorypb.Category, error) {
	resp, err := cl.client.CreateCategory(ctx, &categorypb.CreateCategoryRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return resp.GetCategory(), nil
}

func (cl *CategoryClient) Update(ctx context.Context, number int, req models.UpdateCategoryRequest) (*categorypb.Category, error) {
	resp, err := cl.client.UpdateCategory(ctx, &categorypb.UpdateCategoryRequest{
		Number: int32(number),
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCategory(), nil
}

func (cl *CategoryClient) Delete(ctx context.Context, number int) error {
	_, err := cl.client.DeleteCategory(ctx, &categorypb.DeleteCategoryRequest{Number: int32(number)})
	return err
}

func (cl *CategoryClient) GetByNumber(ctx context.Context, number int) (*categorypb.Category, error) {
	resp, err := cl.client.GetCategory(ctx, &categorypb.GetCategoryRequest{Number: int32(number)})
	if err != nil {
		return nil, err
	}

	return resp.GetCategory(), nil
}

func (cl *CategoryClient) GetAll(ctx context.Context) ([]*categorypb.Category, error) {
	resp, err := cl.client.GetAllCategories(ctx, &categorypb.GetAllCategoriesRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetCategories(), nil
}
