package service

import (
	"context"
	"net/http"

	"github.com/weekendprojectapp/authful/servertools"
	"github.com/weekendprojectapp/authful/signin/internal/providers"
	"github.com/weekendprojectapp/authful/signin/pkg/models"
)

func IsValidUsernamePassword(ctx context.Context, username string, password string) (bool, models.SigninJwt, error) {
	if len(username) == 0 {
		return false, models.SigninJwt{}, servertools.NewServiceError(http.StatusBadRequest, "username is required")
	}

	if len(password) == 0 {
		return false, models.SigninJwt{}, servertools.NewServiceError(http.StatusBadRequest, "password is required")
	}

	return providers.IsValidUsernamePassword(ctx, username, password)
}
