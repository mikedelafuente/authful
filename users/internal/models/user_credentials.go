package models

type UserCredentials struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
