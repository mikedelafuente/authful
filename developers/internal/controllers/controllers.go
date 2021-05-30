package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
	"github.com/mikedelafuente/authful/developers/internal/logger"
	"github.com/mikedelafuente/authful/developers/internal/models"
	"github.com/mikedelafuente/authful/developers/internal/services"
)

func DeveloperSignupPost(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.DeveloperSignupRequest
	json.NewDecoder(r.Body).Decode(&signupRequest)

	userId := r.Context().Value(customclaims.ContextKeyUserId).(string)
	if len(userId) == 0 {
		httptools.HandleError(customerrors.NewServiceError(http.StatusBadRequest, "invalid jwt - user_id is missing"), w)
		return
	}

	// Make sure the username is unique
	user, err := services.CreateDeveloper(r.Context(), userId, signupRequest.OrganizationName, signupRequest.ContactEmail, signupRequest.AgreeToTermsOfService)
	if err != nil {
		logger.Println(err)
		httptools.HandleError(err, w)
		return
	}

	httptools.ProcessResponse(user, w, http.StatusOK)
}

func DevelopersGet(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetDevelopers(r.Context())
	if err != nil {
		logger.Println(err)
		httptools.HandleError(err, w)
		return
	}

	httptools.ProcessResponse(users, w, http.StatusOK)
}
