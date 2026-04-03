package product

type Product struct {
	IdProduct       int     `json:"id_product"`
	CategoryNumber  int     `json:"category_number"`
	ProductName     string  `json:"product_name"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics string  `json:"characteristics"`
}

type ProductCreateRequest struct {
	CategoryNumber  int     `json:"category_number"`
	ProductName     string  `json:"product_name"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics string  `json:"characteristics"`
}
