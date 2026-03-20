package customercard

type CustomerCard struct {
	CardNumber      string  `json:"card_number"`
	CustSurname     string  `json:"cust_surname"`
	CustName        string  `json:"cust_name"`
	CustPatronymic  *string `json:"cust_patronymic,omitempty"`
	PhoneNumber     string  `json:"phone_number"`
	City            *string `json:"city,omitempty"`
	Street          *string `json:"street,omitempty"`
	ZipCode         *string `json:"zip_code,omitempty"`
	Percent         int     `json:"percent"`
}

type CustomerCardCreateRequest struct {
	CardNumber      string  `json:"card_number"`
	CustSurname     string  `json:"cust_surname"`
	CustName        string  `json:"cust_name"`
	CustPatronymic  *string `json:"cust_patronymic,omitempty"`
	PhoneNumber     string  `json:"phone_number"`
	City            *string `json:"city,omitempty"`
	Street          *string `json:"street,omitempty"`
	ZipCode         *string `json:"zip_code,omitempty"`
	Percent         int     `json:"percent"`
}
