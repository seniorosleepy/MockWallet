package models

type User struct {
	ID       int    `json:id,omitempty`
	Username string `json:username`
	Password string `json:password,omitempty`
	Email    string `json:email`
}
