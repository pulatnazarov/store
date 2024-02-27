package models

type IncomeProduct struct {
	ID        string `json:"id"`
	IncomeID  string `json:"income_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}

type CreateIncomeProduct struct {
	IncomeID  string `json:"income_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}

type CreateIncomeProducts struct {
	IncomeProducts []CreateIncomeProduct `json:"income_products"`
}

type IncomeProductsResponse struct {
	IncomeProducts []IncomeProduct
	Count          int
}

type UpdateIncomeProducts struct {
	IncomeProducts []IncomeProduct
}

type IncomeProducts struct {
	IncomeProducts []IncomeProduct
}

type DeleteIncomeProducts struct {
	IDs []PrimaryKey
}

type IncomeProductReport struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	TotalPrice  int    `json:"total_price"`
}

type IncomeProductReportList struct {
	IncomeProducts []IncomeProductReport `json:"income_products"`
	OverallPrice   int                   `json:"overall_price"`
}
