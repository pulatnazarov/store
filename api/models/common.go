package models

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	BasketID string `json:"basket_id"`
	UserID   string `json:"user_id"`
}

type ReportRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	BranchID string `json:"branch_id"`
	From     string `json:"from"`
	To       string `json:"to"`
}
