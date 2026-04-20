package employee

import "time"

type Employee struct {
	ID          string    `json:"id"`
	Surname     string    `json:"surname"`
	Name        string    `json:"name"`
	Patronymic  *string   `json:"patronymic,omitempty"`
	Role        string    `json:"role"`
	Salary      float64   `json:"salary"`
	DateOfBirth time.Time `json:"date_of_birth"`
	DateOfStart time.Time `json:"date_of_start"`
	PhoneNumber string    `json:"phone_number"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	ZipCode     string    `json:"zip_code"`
}

type CreateRequest struct {
	ID          string    `json:"id"`
	Surname     string    `json:"surname"`
	Name        string    `json:"name"`
	Patronymic  *string   `json:"patronymic,omitempty"`
	Role        string    `json:"role"`
	Salary      float64   `json:"salary"`
	DateOfBirth time.Time `json:"date_of_birth"`
	DateOfStart time.Time `json:"date_of_start"`
	PhoneNumber string    `json:"phone_number"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	ZipCode     string    `json:"zip_code"`
}

type CashierPerformance struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Surname        string  `json:"surname"`
	Patronymic     *string `json:"patronymic,omitempty"`
	TotalChecks    int64   `json:"total_checks"`
	TotalItemsSold int64   `json:"total_items_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
}

type BestCashierByPromo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
