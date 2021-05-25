package rest

import (
	"encoding/json"
	"net/http"

	"github.com/weekendprojectapp/authful/server"
	"github.com/weekendprojectapp/authful/users/internal/users/service"
	"github.com/weekendprojectapp/authful/users/pkg/models"
)

func AccountSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if service.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser, err := service.GetUserByUsername(r.Context(), userRequest.Username)
		if err != nil {
			server.HandleError(err, w)
			return
		}
		tokenString, expirationTime, err := service.ProduceJwtTokenForUser(r.Context(), foundUser.Username, foundUser.Id)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			server.HandleError(err, w)
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

		server.ProcessResponse(t, w, http.StatusOK)
	} else {
		server.HandleResponse(w, []byte{}, http.StatusUnauthorized)
	}
}

func AccountSignupPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := service.CreateUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		server.HandleError(err, w)
		return
	}

	server.ProcessResponse(user, w, http.StatusOK)
}

func UsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers(r.Context())
	if err != nil {
		server.HandleError(err, w)
		return
	}

	server.ProcessResponse(users, w, http.StatusOK)
}
