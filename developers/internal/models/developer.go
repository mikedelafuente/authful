package models

import "time"

type Developer struct {
	DeveloperId      string    `json:"developer_id" db:"dev_id"`
	UserId           string    `json:"user_id" db:"user_id"`
	OrganizationName string    `json:"organization_name" db:"organizaiton_name"`
	ContactEmail     string    `json:"contact_email" db:"contact_email"`
	CreateDate       time.Time `json:"create_datetime" db:"create_datetime"`
	UpdateDate       time.Time `json:"update_datetime" db:"update_datetime"`
}
