package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type Balance struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	CurrId int `json:"currency_id"`
	Amount int `json:"amount"`
}
