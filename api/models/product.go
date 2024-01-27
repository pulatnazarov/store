package models

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type CreateProduct struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type UpdateProduct struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type ProductResponse struct {
	Product []Product
	Count   int
}

type ProductSell struct {
	Name          string `json:"name"`
	Price         uint   `json:"-"`
	OriginalPrice uint   `json:"-"`
	Quantity      uint   `json:"quantity"`
}
