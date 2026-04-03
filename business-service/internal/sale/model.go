package sale

type Sale struct {
	Upc           string  `json:"upc"`
	CheckNumber   string  `json:"check_number"`
	ProductNumber int     `json:"product_number"`
	SellingPrice  float64 `json:"selling_price"`
}

type SaleCreateRequest struct {
	Upc           string  `json:"upc"`
	CheckNumber   string  `json:"check_number"`
	ProductNumber int     `json:"product_number"`
	SellingPrice  float64 `json:"selling_price"`
}
