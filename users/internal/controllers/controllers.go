package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/users/internal/models"
	service "github.com/mikedelafuente/authful/users/internal/services"
)

func AccountSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if service.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser, err := service.GetUserByUsername(r.Context(), userRequest.Username)
		if err != nil {
			logger.Error(r.Context(), err)
			httptools.HandleError(r.Context(), err, w)
			return
		}
		tokenString, expirationTime, err := service.ProduceJwtTokenForUser(r.Context(), foundUser.Username, foundUser.UserId)
		if err != nil {
			logger.Error(r.Context(), err)
			// If there is an error in creating the JWT return an internal server error
			httptools.HandleError(r.Context(), err, w)
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

		httptools.ProcessResponse(r.Context(), t, w, http.StatusOK)
	} else {
		httptools.HandleResponse(w, []byte{}, http.StatusUnauthorized)
	}
}

func AccountSignupPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := service.CreateUser(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		logger.Error(r.Context(), err)
		httptools.HandleError(r.Context(), err, w)
		return
	}

	httptools.ProcessResponse(r.Context(), user, w, http.StatusOK)
}

func UsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers(r.Context())
	if err != nil {
		logger.Error(r.Context(), err)
		httptools.HandleError(r.Context(), err, w)
		return
	}

	httptools.ProcessResponse(r.Context(), users, w, http.StatusOK)
}
