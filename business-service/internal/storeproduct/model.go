package storeproduct

type StoreProduct struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	ProductID          int     `json:"product_id"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type DetailedStoreProduct struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
	ProductID          int     `json:"product_id"`
	ProductName        string  `json:"product_name"`
	Producer           *string `json:"producer,omitempty"`
	Characteristics    string  `json:"characteristics"`
	CategoryNumber     int     `json:"category_number"`
	CategoryName       string  `json:"category_name"`
}

type CreateRequest struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	ProductID          int     `json:"product_id"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type CashierSoldAllCategoryProducts struct {
	ID          string  `json:"id"`
	Surname     string  `json:"surname"`
	Name        string  `json:"name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number"`
}
