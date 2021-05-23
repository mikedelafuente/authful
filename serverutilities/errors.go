package serverutilities

import "fmt"

func NewServiceError(httpStatusCode int, error_description string) *ServiceError {
	return &ServiceError{
		Description: error_description,
		StatusCode:  httpStatusCode,
	}
}

type ServiceError struct {
	Description string
	StatusCode  int
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%v: %s", e.StatusCode, e.Description)
}
