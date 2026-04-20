package category

import (
	"context"
	"fmt"

	categorydb "github.com/handziurdmytro/3JlArODA/business-service/internal/category/sqlc"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *categorydb.Queries
}

func NewCategoryRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: categorydb.New(pool)}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) (*Category, error) {
	row, err := r.queries.CreateCategory(ctx, req.Name)
	if err != nil {
		return nil, common.MapRepositoryError("create category", err)
	}

	return mapCategory(row), nil
}

func (r *Repository) Update(ctx context.Context, category Category) (*Category, error) {
	row, err := r.queries.UpdateCategory(ctx, categorydb.UpdateCategoryParams{
		CategoryNumber: int32(category.Number),
		CategoryName:   category.Name,
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("update category %d", category.Number), err)
	}

	return mapCategory(row), nil
}

func (r *Repository) Delete(ctx context.Context, number int) error {
	if _, err := r.GetByNumber(ctx, number); err != nil {
		return err
	}

	if err := r.queries.DeleteCategory(ctx, int32(number)); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete category %d", number), err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Category, error) {
	rows, err := r.queries.GetAllCategoriesSortedByName(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all categories", err)
	}

	categories := make([]Category, 0, len(rows))
	for _, row := range rows {
		categories = append(categories, *mapCategory(row))
	}

	return categories, nil
}

func (r *Repository) GetByNumber(ctx context.Context, number int) (*Category, error) {
	row, err := r.queries.GetCategoryByID(ctx, int32(number))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get category by number %d", number), err)
	}

	return mapCategory(row), nil
}

func (r *Repository) GetStockSummary(ctx context.Context) ([]StockSummary, error) {
	rows, err := r.queries.GetCategoryStockSummary(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get category stock summary", err)
	}

	summaries := make([]StockSummary, 0, len(rows))
	for _, row := range rows {
		avgPrice, err := common.Float64FromNumeric(row.AvgPrice)
		if err != nil {
			return nil, fmt.Errorf("map category stock summary average price: %w", err)
		}

		summaries = append(summaries, StockSummary{
			Number:        int(row.CategoryNumber),
			Name:          row.CategoryName,
			TotalQuantity: row.TotalQuantity,
			AvgPrice:      avgPrice,
		})
	}

	return summaries, nil
}

func mapCategory(row categorydb.Category) *Category {
	return &Category{
		Number: int(row.CategoryNumber),
		Name:   row.CategoryName,
	}
}
