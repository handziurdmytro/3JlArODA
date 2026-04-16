package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	productdb "github.com/handziurdmytro/3JlArODA/business-service/internal/product/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
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
		Producer:        textFromPtr(req.Producer),
		Characteristics: req.Characteristics,
	})
	if err != nil {
		return nil, handleError("create product", err)
	}

	return mapProduct(row), nil
}

func (r *Repository) Update(ctx context.Context, product Product) (*Product, error) {
	row, err := r.queries.UpdateProduct(ctx, productdb.UpdateProductParams{
		IDProduct:       int32(product.ID),
		CategoryNumber:  int32(product.CategoryNumber),
		ProductName:     product.Name,
		Producer:        textFromPtr(product.Producer),
		Characteristics: product.Characteristics,
	})
	if err != nil {
		return nil, handleError(fmt.Sprintf("update product %d", product.ID), err)
	}

	return mapProduct(row), nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	if _, err := r.GetByID(ctx, id); err != nil {
		return err
	}

	if err := r.queries.DeleteProduct(ctx, int32(id)); err != nil {
		return handleError(fmt.Sprintf("delete product %d", id), err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := r.queries.GetAllProductsSortedByName(ctx)
	if err != nil {
		return nil, handleError("get all products", err)
	}

	return mapProducts(rows), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*Product, error) {
	row, err := r.queries.GetProductById(ctx, int32(id))
	if err != nil {
		return nil, handleError(fmt.Sprintf("get product by id %d", id), err)
	}

	return mapProduct(row), nil
}

func (r *Repository) GetByCategory(ctx context.Context, categoryNumber int) ([]Product, error) {
	rows, err := r.queries.GetProductsByCategory(ctx, int32(categoryNumber))
	if err != nil {
		return nil, handleError(fmt.Sprintf("get products by category %d", categoryNumber), err)
	}

	return mapProducts(rows), nil
}

func (r *Repository) SearchByName(ctx context.Context, name string) ([]Product, error) {
	rows, err := r.queries.SearchProductsByName(ctx, textFromString(name))
	if err != nil {
		return nil, handleError(fmt.Sprintf("search products by name %q", name), err)
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
		Producer:        ptrFromText(row.Producer),
		Characteristics: row.Characteristics,
	}
}

func textFromString(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func textFromPtr(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}

	return textFromString(*value)
}

func ptrFromText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func handleError(operation string, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%s: %w", operation, common.ErrNotFound)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return fmt.Errorf("%s: %w", operation, common.ErrAlreadyExists)
		case "23503":
			return fmt.Errorf("%s: %w", operation, common.ErrForeignKeyViolation)
		case "23514":
			return fmt.Errorf("%s: %w", operation, common.ErrCheckViolation)
		}
	}

	return fmt.Errorf("%s: %w", operation, err)
}
