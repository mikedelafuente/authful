package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/signin/internal/models"
	"github.com/mikedelafuente/authful/signin/internal/services"
)

func ApiAccountSigninPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.SigninCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if len(userRequest.Username) == 0 || len(userRequest.Password) == 0 {
		httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusBadRequest, "username and password are required"), w)
		return
	}

	logger.Verbose(r.Context(), "Logging in through UI")
	validLogin, jwt, err := services.IsValidUsernamePassword(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		logger.Error(r.Context(), err)
		httptools.HandleError(r.Context(), err, w)
		return
	} else if !validLogin {
		httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusUnauthorized, "authentication failed"), w)
		return
	}

	httptools.ProcessResponse(r.Context(), jwt, w, http.StatusOK)
}

func ApiAccountSignupPost(w http.ResponseWriter, r *http.Request) {
	var userRequest models.SigninCredentials
	json.NewDecoder(r.Body).Decode(&userRequest)

	if len(userRequest.Username) == 0 || len(userRequest.Password) == 0 {
		httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusBadRequest, "username and password are required"), w)
		return
	}

	logger.Verbose(r.Context(), "Sign up through UI")
	user, err := services.Signup(r.Context(), userRequest.Username, userRequest.Password)
	if err != nil {
		logger.Error(r.Context(), err)
		httptools.HandleError(r.Context(), err, w)
		return
	} else if len(user.UserId) == 0 {
		httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusMethodNotAllowed, "authentication failed"), w)
		return
	}

	httptools.ProcessResponse(r.Context(), user, w, http.StatusOK)
}

func ApiAccountResetPost(w http.ResponseWriter, r *http.Request) {
	var resetRequest models.AccountResetRequest
	json.NewDecoder(r.Body).Decode(&resetRequest)

	httptools.HandleError(r.Context(), customerrors.NewServiceError(http.StatusMethodNotAllowed, "NOT IMPLEMENTED"), w)
}
