package storeproduct

import (
	"context"
	"time"
)

type repository interface {
	Create(ctx context.Context, req CreateRequest) (*StoreProduct, error)
	Update(ctx context.Context, product StoreProduct) (*StoreProduct, error)
	Delete(ctx context.Context, upc string) error
	GetByUPC(ctx context.Context, upc string) (*DetailedStoreProduct, error)
	GetAllSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error)
	GetAllSortedByName(ctx context.Context) ([]DetailedStoreProduct, error)
	GetPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error)
	GetPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error)
	GetNonPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error)
	GetNonPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error)
	GetByCategorySortedByName(ctx context.Context, categoryNumber int) ([]DetailedStoreProduct, error)
	GetCashiersWhoSoldAllProductsFromCategory(ctx context.Context, categoryNumber int, from, to time.Time) ([]CashierSoldAllCategoryProducts, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*StoreProduct, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, product StoreProduct) (*StoreProduct, error) {
	return s.repo.Update(ctx, product)
}

func (s *Service) Delete(ctx context.Context, upc string) error {
	return s.repo.Delete(ctx, upc)
}

func (s *Service) GetByUPC(ctx context.Context, upc string) (*DetailedStoreProduct, error) {
	return s.repo.GetByUPC(ctx, upc)
}

func (s *Service) GetAllSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetAllSortedByQuantity(ctx)
}

func (s *Service) GetAllSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetAllSortedByName(ctx)
}

func (s *Service) GetPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetPromoSortedByQuantity(ctx)
}

func (s *Service) GetPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetPromoSortedByName(ctx)
}

func (s *Service) GetNonPromoSortedByQuantity(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetNonPromoSortedByQuantity(ctx)
}

func (s *Service) GetNonPromoSortedByName(ctx context.Context) ([]DetailedStoreProduct, error) {
	return s.repo.GetNonPromoSortedByName(ctx)
}

func (s *Service) GetByCategorySortedByName(ctx context.Context, categoryNumber int) ([]DetailedStoreProduct, error) {
	return s.repo.GetByCategorySortedByName(ctx, categoryNumber)
}

func (s *Service) GetCashiersWhoSoldAllProductsFromCategory(ctx context.Context, categoryNumber int, from, to time.Time) ([]CashierSoldAllCategoryProducts, error) {
	return s.repo.GetCashiersWhoSoldAllProductsFromCategory(ctx, categoryNumber, from, to)
}
