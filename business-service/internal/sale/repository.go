package sale

import (
	"context"
	"fmt"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	saledb "github.com/handziurdmytro/3JlArODA/business-service/internal/sale/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *saledb.Queries
}

func NewSaleRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: saledb.New(pool)}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) error {
	sellingPrice, err := common.NumericFromFloat64(req.SellingPrice)
	if err != nil {
		return fmt.Errorf("create sale for check %q and upc %q: %w", req.CheckNumber, req.UPC, err)
	}

	err = r.queries.CreateSale(ctx, saledb.CreateSaleParams{
		Upc:           req.UPC,
		CheckNumber:   req.CheckNumber,
		ProductNumber: int32(req.ProductNumber),
		SellingPrice:  sellingPrice,
	})
	if err != nil {
		return common.MapRepositoryError(
			fmt.Sprintf("create sale for check %q and upc %q", req.CheckNumber, req.UPC),
			err,
		)
	}

	return nil
}

func (r *Repository) GetProductSoldQuantity(ctx context.Context, productID int, from, to time.Time) (int64, error) {
	quantity, err := r.queries.GetProductSoldQuantity(ctx, saledb.GetProductSoldQuantityParams{
		IDProduct:   int32(productID),
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return 0, common.MapRepositoryError(fmt.Sprintf("get sold quantity for product %d", productID), err)
	}

	return quantity, nil
}
