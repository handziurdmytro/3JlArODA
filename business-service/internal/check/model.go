package check

import "time"

type Check struct {
	CheckNumber string    `json:"check_number"`
	IdEmployee  string    `json:"id_employee"`
	CardNumber  *string   `json:"card_number,omitempty"`
	PrintDate   time.Time `json:"print_date"`
	SumTotal    float64   `json:"sum_total"`
	Vat         float64   `json:"vat"`
}

type CheckCreateRequest struct {
	CheckNumber string    `json:"check_number"`
	IdEmployee  string    `json:"id_employee"`
	CardNumber  *string   `json:"card_number,omitempty"`
	PrintDate   time.Time `json:"print_date"`
	SumTotal    float64   `json:"sum_total"`
	Vat         float64   `json:"vat"`
}
