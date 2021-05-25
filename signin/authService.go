package main

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/serverutils"
)

type authService struct {
	logic authLogic
}

func newAuthService() *authService {
	s := authService{
		logic: *newAuthLogic(),
	}
	return &s
}

func (s *authService) displayLogin(w http.ResponseWriter, r *http.Request) {

}

func (s *authService) authorizeUser(w http.ResponseWriter, r *http.Request) {
	var userRequest signinCredentialsDto
	json.NewDecoder(r.Body).Decode(&userRequest)

	serverutils.HandleResponse(w, []byte{}, http.StatusUnauthorized)

}
