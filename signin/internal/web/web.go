package web

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/mikedelafuente/authful/servertools"
	"github.com/mikedelafuente/authful/signin/internal/service"
)

type Student struct {
	Name       string
	College    string
	RollNumber int
}

type loginBag struct {
	ErrorMessage string
	Username     string
	FailedLogin  bool
}

type signupBag struct {
	ErrorMessage string
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	bag := loginBag{
		ErrorMessage: "",
		Username:     r.FormValue("username"),
		FailedLogin:  false,
	}

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
		FailedLogin:  false,
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
	bag.FailedLogin = true
	parsedTemplate, _ := template.ParseFiles("template/login.html")
	err = parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

}

func DisplaySignup(w http.ResponseWriter, r *http.Request) {
	bag := loginBag{
		ErrorMessage: "",
		Username:     "",
		FailedLogin:  false,
	}

	parsedTemplate, _ := template.ParseFiles("template/register.html")
	err := parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func ProcessSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	bag := signupBag{
		ErrorMessage: "",
	}

	user, err := service.Signup(r.Context(), username, password)

	if err != nil {
		bag.ErrorMessage = err.Error()
	}

	if len(user.Id) == 0 {
		bag.ErrorMessage = "unable to register"
	} else {
		http.Redirect(w, r, "/login?username="+user.Username, http.StatusFound)
		return
	}

	parsedTemplate, _ := template.ParseFiles("template/register.html")
	err = parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
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
