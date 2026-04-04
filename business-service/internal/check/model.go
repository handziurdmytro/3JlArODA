package check

import "time"

type Check struct {
	Number     string    `json:"number"`
	EmployeeID string    `json:"employee_id"`
	CardNumber *string   `json:"card_number,omitempty"`
	PrintDate  time.Time `json:"print_date"`
	SumTotal   float64   `json:"sum_total"`
	VAT        float64   `json:"vat"`
}

type CreateRequest struct {
	Number     string    `json:"number"`
	EmployeeID string    `json:"employee_id"`
	CardNumber *string   `json:"card_number,omitempty"`
	PrintDate  time.Time `json:"print_date"`
	SumTotal   float64   `json:"sum_total"`
	VAT        float64   `json:"vat"`
}
