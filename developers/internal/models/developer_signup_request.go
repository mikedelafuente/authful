package models

type DeveloperSignupRequest struct {
	AgreeToTermsOfService bool   `json:"agree_to_terms_of_service" db:"agree_to_tos"`
	OrganizationName      string `json:"organization_name" db:"organization_name"`
	ContactEmail          string `json:"contact_email" db:"contact_email"`
}
