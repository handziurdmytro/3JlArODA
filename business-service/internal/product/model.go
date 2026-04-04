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
