package domain

type FullNameRequest struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}
