package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/signin/internal/services"
)

type loginBag struct {
	ErrorMessages []string
	Username      string
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	bag := loginBag{
		ErrorMessages: []string{},
		Username:      r.FormValue("username"),
	}

	parsedTemplate, _ := template.ParseFiles("Templates/login.html")
	err := parsedTemplate.Execute(w, bag)
	if err != nil {
		logger.Error(r.Context(), fmt.Sprintf("Error executing template: %s", err))
		return
	}
}

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	redirectUri := r.FormValue("redirect_uri")

	username := r.FormValue("username")
	password := r.FormValue("password")

	bag := loginBag{
		ErrorMessages: []string{},
		Username:      username,
	}
	logger.Verbose(r.Context(), "Logging in")
	validLogin, jwt, err := services.IsValidUsernamePassword(r.Context(), username, password)
	if err != nil {
		logger.Error(r.Context(), err)
		bag.ErrorMessages = append(bag.ErrorMessages, err.Error())
	} else if !validLogin {
		bag.ErrorMessages = append(bag.ErrorMessages, "authentication failed")
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
	parsedTemplate, _ := template.ParseFiles("Templates/login.html")
	err = parsedTemplate.Execute(w, bag)
	if err != nil {
		logger.Error(r.Context(), fmt.Sprintf("Error executing template: %s", err))
		return
	}
}
