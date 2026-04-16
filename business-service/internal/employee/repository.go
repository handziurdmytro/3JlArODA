package employee

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

type ContactData struct {
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Street      string `json:"street"`
	ZipCode     string `json:"zip_code"`
}

func NewEmployeeRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) CreateEmployee(ctx context.Context, req CreateRequest) error {
	panic("not implemented")
}

func (r *Repository) UpdateEmployeeByID(ctx context.Context, employee Employee) error {
	panic("not implemented")
}

func (r *Repository) DeleteEmployeeByID(ctx context.Context, id string) error {
	panic("not implemented")
}

func (r *Repository) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	panic("not implemented")
}

func (r *Repository) GetEmployeeByID(ctx context.Context, id string) (*Employee, error) {
	panic("not implemented")
}

func (r *Repository) GetEmployeesByRole(ctx context.Context, role string) ([]Employee, error) {
	panic("not implemented")
}

func (r *Repository) GetEmployeeDataBySurname(ctx context.Context, surname string) ([]ContactData, error) {
	panic("not implemented")
}

func (r *Repository) GetEmployeeDataByFullName(ctx context.Context, surname, name, patronymic string) ([]ContactData, error) {
	panic("not implemented")
}
