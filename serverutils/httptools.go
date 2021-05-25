package serverutils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func HandleError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	if e, ok := err.(*ServiceError); ok {
		statusCode = e.StatusCode
	}
	resp := errorResponse{Error: err.Error()}
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

func ProcessResponse(v interface{}, w http.ResponseWriter, statusCode int) {
	b, err := MarshalFormat(v)
	if err != nil {
		// Handle as a server error?
		HandleError(err, w)
		return
	}

	HandleResponse(w, b, statusCode)
}
