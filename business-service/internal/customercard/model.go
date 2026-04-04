package customercard

type CustomerCard struct {
	CardNumber  string  `json:"card_number"`
	Surname     string  `json:"surname"`
	Name        string  `json:"name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
	ZipCode     *string `json:"zip_code,omitempty"`
	Percent     int     `json:"percent"`
}

type CreateRequest struct {
	CardNumber  string  `json:"card_number"`
	Surname     string  `json:"surname"`
	Name        string  `json:"name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
	ZipCode     *string `json:"zip_code,omitempty"`
	Percent     int     `json:"percent"`
}
