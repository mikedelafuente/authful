package httptools

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mikedelafuente/authful/servertools/pkg/customerrors"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ExtractErrorMessageFromJsonBytes(data []byte, defaultMessage string) string {
	if len(data) == 0 {
		return defaultMessage
	}

	var e ErrorResponse
	body := string(data)
	log.Printf("Body:\n%s\n", body)
	err := json.Unmarshal(data, &e)
	if err != nil {
		log.Println(err)
	} else {
		return e.Error
	}

	return defaultMessage

}
func HandleError(err error, w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	if e, ok := err.(*customerrors.ServiceError); ok {
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

func ProcessResponse(v interface{}, w http.ResponseWriter, statusCode int) {
	b, err := MarshalFormat(v)
	if err != nil {
		// Handle as a server error?
		HandleError(err, w)
		return
	}

	HandleResponse(w, b, statusCode)
}

// Validates that the response code is a 2xx
func IsOkResponse(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
