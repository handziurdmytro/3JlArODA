package product

type Product struct {
	ID              int     `json:"id"`
	CategoryNumber  int     `json:"category_number"`
	Name            string  `json:"name"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics string  `json:"characteristics"`
}

type CreateRequest struct {
	CategoryNumber  int     `json:"category_number"`
	Name            string  `json:"name"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics string  `json:"characteristics"`
}

type SalesStats struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Producer          *string `json:"producer,omitempty"`
	Characteristics   string  `json:"characteristics"`
	CategoryNumber    int     `json:"category_number"`
	CategoryName      string  `json:"category_name"`
	TotalSoldQuantity int64   `json:"total_sold_quantity"`
	TotalRevenue      float64 `json:"total_revenue"`
}
