package serverutils

import (
	"fmt"
	"net/http"
)

func GenerateHtmlHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html, charset=UTF-8")

	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html lang=\"en\">")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
}

func GenerateHtmlFooter(w http.ResponseWriter) {
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")

}
