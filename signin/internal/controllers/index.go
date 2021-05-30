package controllers

import (
	"html/template"
	"net/http"

	"github.com/mikedelafuente/authful/signin/internal/logger"
)

type Student struct {
	Name       string
	College    string
	RollNumber int
}

func Index(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Name:       "GB",
		College:    "GolangBlogs",
		RollNumber: 1,
	}
	parsedTemplate, _ := template.ParseFiles("Templates/index.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		logger.Println("Error executing template :", err)
		return
	}
}
