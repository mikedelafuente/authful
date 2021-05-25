package rest

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/serverutils"
	svc "github.com/weekendprojectapp/authful/users/internal/users/service"
	"github.com/weekendprojectapp/authful/users/pkg/models"
)

type User struct {
	/// Handles the marshalling and unmarshalling of service requests

	// The business logic for users
	logic svc.User
}

func New() *User {
	s := User{
		logic: *svc.New(),
	}
	return &s
}

func (s *User) AccountSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if s.logic.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser, err := s.logic.GetUserByUsername(r.Context(), userRequest.Username)
		if err != nil {
			serverutils.HandleError(err, w)
			return
		}
		tokenString, expirationTime, err := s.logic.ProduceJwtTokenForUser(r.Context(), foundUser.Username, foundUser.Id)
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

func (s *User) AccountSignupPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := s.logic.CreateUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		serverutils.HandleError(err, w)
		return
	}

	serverutils.ProcessResponse(user, w, http.StatusOK)
}

func (s *User) UsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := s.logic.GetUsers(r.Context())
	if err != nil {
		serverutils.HandleError(err, w)
		return
	}

	serverutils.ProcessResponse(users, w, http.StatusOK)
}
