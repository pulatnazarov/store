package models

type CustomerLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
