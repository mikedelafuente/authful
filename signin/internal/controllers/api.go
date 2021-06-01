package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
	"github.com/mikedelafuente/authful/signin/internal/models"
)

func ApiSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.SigninCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if len(userRequest.Username) == 0 || len(userRequest.Password) == 0 {
		httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusBadRequest, "username and password are required"), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "apiTest",
		Value:    "A test",
		Expires:  time.Now().Add(time.Duration(time.Duration.Minutes(5))),
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// logger.Verbose(r.Context(), "Logging in through UI")
	// validLogin, jwt, err := services.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password)
	// if err != nil {
	// 	logger.Error(r.Context(), err)
	// 	httptools.HandleError(r.Context(), err, w)
	// 	return
	// } else if !validLogin {
	// 	httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusUnauthorized, "authentication failed"), w)
	// 	return
	// }

	// if validLogin {
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:    "userSessionToken",
	// 		Value:   jwt.Jwt,
	// 		Expires: jwt.Expires,
	// 	})
	// }

	httptools.HandleResponse(w, []byte{}, http.StatusOK)
}
