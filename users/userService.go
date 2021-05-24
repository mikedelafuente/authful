package main

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/serverutils"
)

type userService struct {
	/// Handles the marshalling and unmarshalling of service requests

	// The business logic for users
	logic userLogic
}

func newUserService() *userService {
	s := userService{
		logic: *newUserLogic(),
	}
	return &s
}

func (s *userService) authorizeUser(w http.ResponseWriter, r *http.Request) {
	var userRequest userCredentialsDto
	json.NewDecoder(r.Body).Decode(&userRequest)

	if s.logic.isValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser, err := s.logic.getUserByUsername(r.Context(), userRequest.Username)
		if err != nil {
			serverutils.HandleError(err, w)
			return
		}
		tokenString, expirationTime, err := s.logic.produceJwtTokenForUser(r.Context(), foundUser.Username, foundUser.Id)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			serverutils.HandleError(err, w)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		t := map[string]interface{}{}
		t["jwt"] = tokenString
		t["expires"] = expirationTime

		serverutils.ProcessResponse(t, w, http.StatusOK)
	} else {
		serverutils.HandleResponse(w, []byte{}, http.StatusUnauthorized)
	}
}

func (s *userService) createUser(w http.ResponseWriter, r *http.Request) {
	var userRequest userCredentialsDto
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := s.logic.createUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		serverutils.HandleError(err, w)
		return
	}

	serverutils.ProcessResponse(user, w, http.StatusOK)

}

func (s *userService) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.logic.getUsers(r.Context())
	if err != nil {
		serverutils.HandleError(err, w)
		return
	}

	serverutils.ProcessResponse(users, w, http.StatusOK)
}
