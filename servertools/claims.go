package servertools

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	SystemId string `json:"id"`
	Type     string `json:"type"`
	jwt.StandardClaims
}
