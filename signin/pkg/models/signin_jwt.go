package models

import "time"

type SigninJwt struct {
	Expires time.Time `json:"expires"`
	Jwt     string    `json:"jwt"`
}
