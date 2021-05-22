package main

import (
	"encoding/json"
	"log"
	"net/http"

	"weekendproject.app/authful/users/config"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authorize not implemented yet"))
}

func (s *userService) createUser(w http.ResponseWriter, r *http.Request) {
	var userRequest userCreateDto
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := s.logic.createUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	b, _ := json.MarshalIndent(user, config.JsonMarshalPrefix, config.JsonMarshalPrefix)
	w.Write(b)
}

func (s *userService) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.logic.getUsers(r.Context())
	if err != nil {
		// Handle as a server error?
		log.Print(err)
		return
	}

	b, err := json.MarshalIndent(users, config.JsonMarshalPrefix, config.JsonMarshalPrefix)
	if err != nil {
		// Handle as a server error?
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
