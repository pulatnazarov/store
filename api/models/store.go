package models

type Dealer struct{}

type Repository struct {
	ProductList []Product `json:"product_list"`
}

type Store struct {
	Dealer Dealer     `json:"dealer"`
	Repo   Repository `json:"repo"`
	Budget int        `json:"budget"`
	Profit int        `json:"profit"`
}

type Ticket struct {
	Name         string `json:"name"`
	TotalPrice   uint   `json:"total_price"`
	SoldQuantity uint   `json:"sold_quantity"`
}
