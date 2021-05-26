package models

import "time"

type User struct {
	Id         string    `json:"id" db:"id"`
	Username   string    `json:"username" db:"username"`
	CreateDate time.Time `json:"create_datetime" db:"create_datetime"`
	UpdateDate time.Time `json:"update_datetime" db:"update_datetime"`
}
