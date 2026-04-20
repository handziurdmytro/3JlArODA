package category

type Category struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type CreateRequest struct {
	Name string `json:"name"`
}

type StockSummary struct {
	Number        int     `json:"number"`
	Name          string  `json:"name"`
	TotalQuantity int64   `json:"total_quantity"`
	AvgPrice      float64 `json:"avg_price"`
}
