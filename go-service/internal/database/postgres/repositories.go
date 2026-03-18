package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	Employees     *EmployeeRepository
	Categories    *CategoryRepository
	Checks        *CheckRepository
	Products      *ProductRepository
	StoreProducts *StoreProductRepository
	Sales         *SaleRepository
	CustomerCards *CustomerCardRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Employees:     NewEmployeeRepository(pool),
		Categories:    NewCategoryRepository(pool),
		Checks:        NewCheckRepository(pool),
		Products:      NewProductRepository(pool),
		StoreProducts: NewStoreProductRepository(pool),
		Sales:         NewSaleRepository(pool),
		CustomerCards: NewCustomerCardRepository(pool),
	}
}

