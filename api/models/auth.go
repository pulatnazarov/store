package models

type CustomerLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CustomerLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
