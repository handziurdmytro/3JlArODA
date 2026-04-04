package sale

type Sale struct {
	UPC           string  `json:"upc"`
	CheckNumber   string  `json:"check_number"`
	ProductNumber int     `json:"product_number"`
	SellingPrice  float64 `json:"selling_price"`
}

type CreateRequest struct {
	UPC           string  `json:"upc"`
	CheckNumber   string  `json:"check_number"`
	ProductNumber int     `json:"product_number"`
	SellingPrice  float64 `json:"selling_price"`
}
