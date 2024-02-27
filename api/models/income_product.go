package models

type IncomeProduct struct {
	ID        string `json:"id"`
	IncomeID  string `json:"income_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	BranchID  string `json:"branch_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletadAt int    `json:"deleted_at"`
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

//

type IncomeProductReport struct {
	IncomeProdutName string `json:"income_product_name"`
	Quantity         int    `json:"quantity"`
	Price            int    `json:"price"`
	TotalPrice       int    `json:"total_price"`
}
type IncomeProductReportList struct {
	IncomeProducts []IncomeProductReport `json:"income_products_report"`
	OverallPrice   int                   `json:"overall_price"`
	Count          int                   `json:"count"`
}
type IncomeProductReportListRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	BranchID string `json:"branch_id"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}
