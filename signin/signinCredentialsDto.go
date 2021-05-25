package main

type signinCredentialsDto struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
