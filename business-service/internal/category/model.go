package category

type Category struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type CreateRequest struct {
	Name string `json:"name"`
}
