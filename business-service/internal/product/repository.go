package product

import (
	"context"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	productdb "github.com/handziurdmytro/3JlArODA/business-service/internal/product/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *productdb.Queries
}

func NewProductRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		queries: productdb.New(pool),
	}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) (*Product, error) {
	row, err := r.queries.CreateProduct(ctx, productdb.CreateProductParams{
		CategoryNumber:  int32(req.CategoryNumber),
		ProductName:     req.Name,
		Producer:        common.TextFromPtr(req.Producer),
		Characteristics: req.Characteristics,
	})
	if err != nil {
		return nil, common.MapRepositoryError("create product", err)
	}

	return mapProduct(row), nil
}

func (r *Repository) Update(ctx context.Context, product Product) (*Product, error) {
	row, err := r.queries.UpdateProduct(ctx, productdb.UpdateProductParams{
		IDProduct:       int32(product.ID),
		CategoryNumber:  int32(product.CategoryNumber),
		ProductName:     product.Name,
		Producer:        common.TextFromPtr(product.Producer),
		Characteristics: product.Characteristics,
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("update product %d", product.ID), err)
	}

	return mapProduct(row), nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	if _, err := r.GetByID(ctx, id); err != nil {
		return err
	}

	if err := r.queries.DeleteProduct(ctx, int32(id)); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete product %d", id), err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := r.queries.GetAllProductsSortedByName(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all products", err)
	}

	return mapProducts(rows), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*Product, error) {
	row, err := r.queries.GetProductById(ctx, int32(id))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get product by id %d", id), err)
	}

	return mapProduct(row), nil
}

func (r *Repository) GetByCategory(ctx context.Context, categoryNumber int) ([]Product, error) {
	rows, err := r.queries.GetProductsByCategory(ctx, int32(categoryNumber))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get products by category %d", categoryNumber), err)
	}

	return mapProducts(rows), nil
}

func (r *Repository) SearchByName(ctx context.Context, name string) ([]Product, error) {
	rows, err := r.queries.SearchProductsByName(ctx, common.TextFromString(name))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("search products by name %q", name), err)
	}

	return mapProducts(rows), nil
}

func mapProducts(rows []productdb.Product) []Product {
	products := make([]Product, 0, len(rows))
	for _, row := range rows {
		products = append(products, *mapProduct(row))
	}

	return products
}

func mapProduct(row productdb.Product) *Product {
	return &Product{
		ID:              int(row.IDProduct),
		CategoryNumber:  int(row.CategoryNumber),
		Name:            row.ProductName,
		Producer:        common.PtrFromText(row.Producer),
		Characteristics: row.Characteristics,
	}
}
