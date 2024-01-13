package models

type User struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
	UserType string `json:"user_type"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
	UserType string `json:"user_type"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Cash     uint   `json:"cash"`
}

type PrimaryKey struct {
	ID string `json:"id"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int   `json:"count"`
}

type GetListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
