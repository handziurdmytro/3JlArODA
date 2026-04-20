package storeproduct

import (
	"context"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	storeproductdb "github.com/handziurdmytro/3JlArODA/business-service/internal/storeproduct/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *storeproductdb.Queries
}

func NewStoreProductRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: storeproductdb.New(pool)}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) (*StoreProduct, error) {
	sellingPrice, err := common.NumericFromFloat64(req.SellingPrice)
	if err != nil {
		return nil, fmt.Errorf("create store product %q: %w", req.UPC, err)
	}

	row, err := r.queries.CreateStoreProduct(ctx, storeproductdb.CreateStoreProductParams{
		Upc:                req.UPC,
		UpcProm:            common.TextFromPtr(req.UPCProm),
		IDProduct:          int32(req.ProductID),
		SellingPrice:       sellingPrice,
		ProductsNumber:     int32(req.ProductsNumber),
		PromotionalProduct: req.PromotionalProduct,
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("create store product %q", req.UPC), err)
	}

	return mapStoreProduct(row)
}

func (r *Repository) Update(ctx context.Context, product StoreProduct) (*StoreProduct, error) {
	sellingPrice, err := common.NumericFromFloat64(product.SellingPrice)
	if err != nil {
		return nil, fmt.Errorf("update store product %q: %w", product.UPC, err)
	}

	row, err := r.queries.UpdateStoreProduct(ctx, storeproductdb.UpdateStoreProductParams{
		Upc:                product.UPC,
		UpcProm:            common.TextFromPtr(product.UPCProm),
		IDProduct:          int32(product.ProductID),
		SellingPrice:       sellingPrice,
		ProductsNumber:     int32(product.ProductsNumber),
		PromotionalProduct: product.PromotionalProduct,
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("update store product %q", product.UPC), err)
	}

	return mapStoreProduct(row)
}

func (r *Repository) Delete(ctx context.Context, upc string) error {
	if _, err := r.GetByUPC(ctx, upc); err != nil {
		return err
	}

	if err := r.queries.DeleteStoreProduct(ctx, upc); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete store product %q", upc), err)
	}

	return nil
}

func (r *Repository) GetByUPC(ctx context.Context, upc string) (*DetailedStoreProduct, error) {
	row, err := r.queries.GetStoreProductByUPC(ctx, upc)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get store product by upc %q", upc), err)
	}

	return mapDetailedStoreProduct(
		row.Upc,
		row.UpcProm,
		row.SellingPrice,
		row.ProductsNumber,
		row.PromotionalProduct,
		row.IDProduct,
		row.ProductName,
		row.Producer,
		row.Characteristics,
		row.CategoryNumber,
		row.CategoryName,
	)
}

func (r *Repository) GetAllSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetAllStoreProductsSortedByQuantity(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all store products sorted by quantity", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetAllSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetAllStoreProductsSortedByName(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all store products sorted by name", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetPromoStoreProductsSortedByQuantity(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get promo store products sorted by quantity", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetPromoStoreProductsSortedByName(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get promo store products sorted by name", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetNonPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetNonPromoStoreProductsSortedByQuantity(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get non-promo store products sorted by quantity", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetNonPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetNonPromoStoreProductsSortedByName(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get non-promo store products sorted by name", err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repository) GetByCategorySortedByName(ctx context.Context, categoryNumber int) ([]DetailedStoreProduct, error) {
	rows, err := r.queries.GetStoreProductsByCategorySortedByName(ctx, int32(categoryNumber))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get store products by category %d sorted by name", categoryNumber), err)
	}

	result := make([]DetailedStoreProduct, 0, len(rows))
	for _, row := range rows {
		product, err := mapDetailedStoreProduct(
			row.Upc,
			row.UpcProm,
			row.SellingPrice,
			row.ProductsNumber,
			row.PromotionalProduct,
			row.IDProduct,
			row.ProductName,
			row.Producer,
			row.Characteristics,
			row.CategoryNumber,
			row.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func mapStoreProduct(row storeproductdb.StoreProduct) (*StoreProduct, error) {
	sellingPrice, err := common.Float64FromNumeric(row.SellingPrice)
	if err != nil {
		return nil, fmt.Errorf("map store product selling price: %w", err)
	}

	return &StoreProduct{
		UPC:                row.Upc,
		UPCProm:            common.PtrFromText(row.UpcProm),
		ProductID:          int(row.IDProduct),
		SellingPrice:       sellingPrice,
		ProductsNumber:     int(row.ProductsNumber),
		PromotionalProduct: row.PromotionalProduct,
	}, nil
}

func mapDetailedStoreProduct(
	upc string,
	upcProm pgtype.Text,
	sellingPriceValue pgtype.Numeric,
	productsNumber int32,
	promotionalProduct bool,
	productID int32,
	productName string,
	producer pgtype.Text,
	characteristics string,
	categoryNumber int32,
	categoryName string,
) (*DetailedStoreProduct, error) {
	sellingPrice, err := common.Float64FromNumeric(sellingPriceValue)
	if err != nil {
		return nil, fmt.Errorf("map detailed store product selling price: %w", err)
	}

	return &DetailedStoreProduct{
		UPC:                upc,
		UPCProm:            common.PtrFromText(upcProm),
		SellingPrice:       sellingPrice,
		ProductsNumber:     int(productsNumber),
		PromotionalProduct: promotionalProduct,
		ProductID:          int(productID),
		ProductName:        productName,
		Producer:           common.PtrFromText(producer),
		Characteristics:    characteristics,
		CategoryNumber:     int(categoryNumber),
		CategoryName:       categoryName,
	}, nil
}
