package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikedelafuente/authful/developers/internal/models"
	"github.com/mikedelafuente/authful/developers/internal/services"
	"github.com/mikedelafuente/authful/servertools/pkg/httptools"
)

func DeveloperSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.DeveloperCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if services.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password) {
		// Generate a JWT and return it
		foundUser, err := services.GetDeveloperByUsername(r.Context(), userRequest.Username)
		if err != nil {
			httptools.HandleError(err, w)
			return
		}
		tokenString, expirationTime, err := services.ProduceJwtTokenForDeveloper(r.Context(), foundUser.Username, foundUser.Id)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			httptools.HandleError(err, w)
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

		httptools.ProcessResponse(t, w, http.StatusOK)
	} else {
		httptools.HandleResponse(w, []byte{}, http.StatusUnauthorized)
	}
}

func DeveloperSignupPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.DeveloperCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	// Make sure the username is unique
	user, err := services.CreateDeveloper(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		httptools.HandleError(err, w)
		return
	}

	httptools.ProcessResponse(user, w, http.StatusOK)
}

func DevelopersGet(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetDevelopers(r.Context())
	if err != nil {
		httptools.HandleError(err, w)
		return
	}

	httptools.ProcessResponse(users, w, http.StatusOK)
}
