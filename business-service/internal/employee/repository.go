package employee

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (repository Repository) CreateEmployee(ctx context.Context, employee Employee) error {
	_, err := repository.pool.Exec(ctx,
		`INSERT INTO employees (
			id_employee, empl_surname, empl_name, empl_patronymic, 
			empl_role, salary, date_of_birth, date_of_start, 
			phone_number, city, street, zip_code
		) VALUES (
			$1, $2, $3, $4, 
			$5, $6, $7, $8, 
			$9, $10, $11, $12
		)`,
		employee.ID, employee.Surname, employee.Name, employee.Patronymic,
		employee.Role, employee.Salary, employee.DateOfBirth, employee.DateOfStart,
		employee.PhoneNumber, employee.City, employee.Street, employee.ZipCode,
	)
	return err
}
