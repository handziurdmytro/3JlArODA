package storeproduct

type StoreProduct struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	ProductID          int     `json:"product_id"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type CreateRequest struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	ProductID          int     `json:"product_id"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}
