package models

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
	BranchID      string `json:"branch_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletadAt     int    `json:"deleted_at"`
}

type CreateProduct struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
	BranchID      string `json:"branch_id"`
}

type UpdateProduct struct {
	ID            string `json:"-"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type ProductResponse struct {
	Products []Product
	Count    int
}

type ProductSell struct {
	SelectedProducts       SellRequest    `json:"selected_products"`
	ProductPrices          map[string]int `json:"product_prices"`
	NotEnoughProducts      map[string]int `json:"not_enough_products"`
	NotEnoughProductPrices map[string]int `json:"prices"`
	ProductsBranchID       string         `json:"products_branch_id"`
	Check                  Check          `json:"check"`
	OnlyOneProductIDs      []string       `json:"only_one_product_ids"`
}

type SellRequest struct {
	Products map[string]int `json:"products"`
	BasketID string         `json:"basket_id"`
	BranchID string         `json:"branch_id"`
}

type DeliverProducts struct {
	NotEnoughProducts map[string]int `json:"not_enough_products"`
	NewProducts       map[string]int `json:"new_products"`
	NewProductPrices  map[string]int `json:"new_product_prices"`
}

type Check struct {
	Products []Product `json:"products"`
	TotalSum int       `json:"total_sum"`
}

type ProductCase struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
	BranchID      string `json:"branch_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Want          interface{}
}

// //////// product report
type ProductReport struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	TotalPrice  int    `json:"total_price"`
}

type ProductReportList struct {
	Products     []ProductReport `json:"products"`
	OverallPrice int             `json:"overall_price"`
	Count        int             `json:"count"`
}

type ProductReportListRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	BranchId string `json:"branch_id"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}
