package customclaims

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"user_id"`
	jwt.StandardClaims
}
