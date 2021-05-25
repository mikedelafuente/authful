package web

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/server"
	"github.com/weekendprojectapp/authful/signin/pkg/models"
)

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	server.HandleResponse(w, []byte{}, http.StatusOK)
}

func AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.SigninCredentials
	json.NewDecoder(r.Body).Decode(&loginRequest)

	server.HandleResponse(w, []byte{}, http.StatusUnauthorized)

}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	server.HandleResponse(w, []byte{}, http.StatusOK)
}
