package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/weekendprojectapp/authful/servertools"
	"github.com/weekendprojectapp/authful/signin/pkg/models"
)

type Student struct {
	Name       string
	College    string
	RollNumber int
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("template/login.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.SigninCredentials
	json.NewDecoder(r.Body).Decode(&loginRequest)

	servertools.HandleResponse(w, []byte{}, http.StatusUnauthorized)

}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	servertools.HandleResponse(w, []byte{}, http.StatusOK)
}

func Index(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Name:       "GB",
		College:    "GolangBlogs",
		RollNumber: 1,
	}
	parsedTemplate, _ := template.ParseFiles("Template/index.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
