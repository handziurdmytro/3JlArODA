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

type FullCheckData struct {
	CheckNumber        string    `json:"check_number"`
	PrintDate          time.Time `json:"print_date"`
	SumTotal           float64   `json:"sum_total"`
	VAT                float64   `json:"vat"`
	UPC                string    `json:"upc"`
	Quantity           int       `json:"quantity"`
	SellingPrice       float64   `json:"selling_price"`
	ProductName        string    `json:"product_name"`
	EmployeeSurname    string    `json:"employee_surname"`
	EmployeeName       string    `json:"employee_name"`
	EmployeePatronymic *string   `json:"employee_patronymic,omitempty"`
}

type Detail struct {
	CheckNumber  string    `json:"check_number"`
	PrintDate    time.Time `json:"print_date"`
	SumTotal     float64   `json:"sum_total"`
	ProductName  string    `json:"product_name"`
	UPC          string    `json:"upc"`
	Quantity     int       `json:"quantity"`
	SellingPrice float64   `json:"selling_price"`
}
