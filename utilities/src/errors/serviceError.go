package errors

func NewServiceError(httpStatusCode int, error_description string) ServiceError {
	return ServiceError{
		Error:      error_description,
		StatusCode: httpStatusCode,
	}
}

type ServiceError struct {
	Error      string
	StatusCode int
}
