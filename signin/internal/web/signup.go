package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mikedelafuente/authful/signin/internal/service"
)

type signupBag struct {
	Blah          []int
	ErrorMessages []string
	Username      string
}

func DisplaySignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	bag := signupBag{
		Username:      username,
		ErrorMessages: []string{},
		Blah:          []int{},
	}

	//bag.ErrorMessages = servertools.ConvertLineBreaksToHtml(bag.ErrorMessage)
	parsedTemplate, _ := template.ParseFiles("template/signup.html")
	err := parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func ProcessSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	verifyPassword := r.FormValue("verify-password")

	bag := signupBag{
		Username:      username,
		ErrorMessages: []string{},
	}

	if len(username) == 0 {

		bag.ErrorMessages = append(bag.ErrorMessages, "Enter a username")

	}

	if len(password) == 0 {
		bag.ErrorMessages = append(bag.ErrorMessages, "Enter a password.")
	} else {
		if password != verifyPassword {
			bag.ErrorMessages = append(bag.ErrorMessages, "Please verify the password. Passwords do not match.")
		}
	}

	if len(bag.ErrorMessages) == 0 {

		user, err := service.Signup(r.Context(), username, password)

		if err != nil {
			bag.ErrorMessages = append(bag.ErrorMessages, err.Error())
		}

		if len(user.Id) == 0 {
			bag.ErrorMessages = append(bag.ErrorMessages, "Unable to register for an account")
		} else {
			http.Redirect(w, r, "/login?username="+user.Username, http.StatusFound)
			return
		}
	}

	parsedTemplate, _ := template.ParseFiles("template/signup.html")
	err := parsedTemplate.Execute(w, bag)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

}
