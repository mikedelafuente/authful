package web

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/mikedelafuente/authful/servertools"
	"github.com/mikedelafuente/authful/signin/internal/service"
)

type loginBag struct {
	ErrorMessage string
	Username     string
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	bag := loginBag{
		ErrorMessage: "",
		Username:     r.FormValue("username"),
	}

	bag.ErrorMessage = servertools.ConvertLineBreaksToHtml(bag.ErrorMessage)
	parsedTemplate, _ := template.ParseFiles("template/login.html")
	err := parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	redirectUri := r.FormValue("redirect_uri")

	username := r.FormValue("username")
	password := r.FormValue("password")

	bag := loginBag{
		ErrorMessage: "",
		Username:     username,
	}

	validLogin, jwt, err := service.IsValidUsernamePassword(r.Context(), username, password)
	if err != nil {
		bag.ErrorMessage = err.Error()
	}

	if validLogin {
		http.SetCookie(w, &http.Cookie{
			Name:    "userSessionToken",
			Value:   jwt.Jwt,
			Expires: jwt.Expires,
		})

		if len(redirectUri) > 0 {
			redirectUri, _ = url.QueryUnescape(redirectUri)
			http.Redirect(w, r, redirectUri, http.StatusFound)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		return
	}

	bag.Username = username
	bag.ErrorMessage = servertools.ConvertLineBreaksToHtml(bag.ErrorMessage)
	parsedTemplate, _ := template.ParseFiles("template/login.html")
	err = parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

}
