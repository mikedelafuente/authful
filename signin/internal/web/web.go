package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/weekendprojectapp/authful/servertools"
	"github.com/weekendprojectapp/authful/signin/internal/service"
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
	//redirectUri := r.FormValue("redirect_uri")

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Authorizing user")
	validLogin, jwt, err := service.IsValidUsernamePassword(r.Context(), username, password)
	if err != nil {
		servertools.HandleError(err, w)
		return
	}

	if validLogin {
		servertools.ProcessResponse(jwt, w, http.StatusOK)
	} else {
		servertools.HandleResponse(w, []byte("bad credentials"), http.StatusUnauthorized)
	}

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
