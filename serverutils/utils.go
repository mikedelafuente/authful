package serverutils

import (
	"fmt"
	"net/http"
)

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

func HandleError(err error, w http.ResponseWriter) {
	if e, ok := err.(*ServiceError); ok {
		w.WriteHeader(e.StatusCode)
		w.Write([]byte(e.Description))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
