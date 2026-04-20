package category

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	categorypb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/category"
)

type GRPCHandler struct {
	categorypb.UnimplementedCategoryServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateCategory(ctx context.Context, req *categorypb.CreateCategoryRequest) (*categorypb.CategoryResponse, error) {
	category, err := h.service.Create(ctx, CreateRequest{Name: req.GetName()})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.CategoryResponse{Category: toProtoCategory(category)}, nil
}

func (h *GRPCHandler) UpdateCategory(ctx context.Context, req *categorypb.UpdateCategoryRequest) (*categorypb.CategoryResponse, error) {
	category, err := h.service.Update(ctx, Category{
		Number: int(req.GetNumber()),
		Name:   req.GetName(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.CategoryResponse{Category: toProtoCategory(category)}, nil
}

func (h *GRPCHandler) DeleteCategory(ctx context.Context, req *categorypb.DeleteCategoryRequest) (*categorypb.DeleteCategoryResponse, error) {
	if err := h.service.Delete(ctx, int(req.GetNumber())); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.DeleteCategoryResponse{Success: true}, nil
}

func (h *GRPCHandler) GetCategory(ctx context.Context, req *categorypb.GetCategoryRequest) (*categorypb.CategoryResponse, error) {
	category, err := h.service.GetByNumber(ctx, int(req.GetNumber()))
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.CategoryResponse{Category: toProtoCategory(category)}, nil
}

func (h *GRPCHandler) GetAllCategories(ctx context.Context, _ *categorypb.GetAllCategoriesRequest) (*categorypb.GetCategoriesResponse, error) {
	categories, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.GetCategoriesResponse{Categories: toProtoCategories(categories)}, nil
}

func (h *GRPCHandler) GetCategoryStockSummary(ctx context.Context, _ *categorypb.GetCategoryStockSummaryRequest) (*categorypb.GetCategoryStockSummaryResponse, error) {
	summaries, err := h.service.GetStockSummary(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &categorypb.GetCategoryStockSummaryResponse{Summaries: toProtoStockSummaries(summaries)}, nil
}

func toProtoCategories(categories []Category) []*categorypb.Category {
	result := make([]*categorypb.Category, 0, len(categories))
	for _, category := range categories {
		category := category
		result = append(result, toProtoCategory(&category))
	}

	return result
}

func toProtoCategory(category *Category) *categorypb.Category {
	if category == nil {
		return nil
	}

	return &categorypb.Category{
		Number: int32(category.Number),
		Name:   category.Name,
	}
}

func toProtoStockSummaries(summaries []StockSummary) []*categorypb.CategoryStockSummary {
	result := make([]*categorypb.CategoryStockSummary, 0, len(summaries))
	for _, summary := range summaries {
		result = append(result, &categorypb.CategoryStockSummary{
			Number:        int32(summary.Number),
			Name:          summary.Name,
			TotalQuantity: summary.TotalQuantity,
			AvgPrice:      summary.AvgPrice,
		})
	}

	return result
}
