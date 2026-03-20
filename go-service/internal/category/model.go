package category

type Category struct {
	CategoryNumber int    `json:"category_number"`
	CategoryName   string `json:"category_name"`
}

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name"`
}
