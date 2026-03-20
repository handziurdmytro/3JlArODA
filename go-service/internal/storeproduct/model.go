package storeproduct

type StoreProduct struct {
	Upc                string  `json:"upc"`
	UpcProm            *string `json:"upc_prom,omitempty"`
	IdProduct          int     `json:"id_product"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type StoreProductCreateRequest struct {
	Upc                string  `json:"upc"`
	UpcProm            *string `json:"upc_prom,omitempty"`
	IdProduct          int     `json:"id_product"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}
