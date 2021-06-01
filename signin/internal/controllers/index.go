package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
)

func Index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("Templates/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		logger.Error(r.Context(), fmt.Sprintf("Error executing template : %s", err))
		return
	}
}
