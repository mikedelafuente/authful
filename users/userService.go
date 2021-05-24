package main

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/serverutils"
	"github.com/weekendprojectapp/authful/users/config"
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

	if s.logic.IsValidUsernamePassword(userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser := s.logic.GetUserByUsername(userRequest.Username)

		tokenString, expirationTime, err := s.logic.ProduceJwtTokenForUser(foundUser.Username, foundUser.Id)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			serverutils.HandleError(err)
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
		b, _ := serverutils.MarshalFormat(t)

		serverutils.HandleResponse(w, b, http.StatusOK)
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

	w.WriteHeader(http.StatusCreated)
	b, _ := json.MarshalIndent(user, config.JsonMarshalPrefix, config.JsonMarshalPrefix)
	w.Write(b)
}

func (s *userService) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.logic.getUsers(r.Context())
	if err != nil {
		serverutils.HandleError(err, w)
		return
	}

	b, err := json.MarshalIndent(users, config.JsonMarshalPrefix, config.JsonMarshalPrefix)
	if err != nil {
		// Handle as a server error?
		serverutils.HandleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
