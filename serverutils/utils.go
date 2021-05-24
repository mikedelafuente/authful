package serverutils

import (
	"encoding/json"
	"net/http"
)

func NewServiceError(httpStatusCode int, error_description string) *ServiceError {
	return &ServiceError{
		Description: error_description,
		StatusCode:  httpStatusCode,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ServiceError struct {
	Description string
	StatusCode  int
}

func (e *ServiceError) Error() string {
	return e.Description
}

func HandleError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	if e, ok := err.(*ServiceError); ok {
		statusCode = e.StatusCode
	}
	resp := ErrorResponse{Error: err.Error()}
	b, _ := MarshalFormat(resp)
	HandleResponse(w, b, statusCode)
}

func HandleResponse(w http.ResponseWriter, b []byte, statusCode int) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(b)
}

func MarshalFormat(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
