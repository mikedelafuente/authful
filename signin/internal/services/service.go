package services

import (
	"context"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful/signin/internal/models"
	"github.com/mikedelafuente/authful/signin/internal/providers"
)

func IsValidUsernamePassword(ctx context.Context, username string, password string) (bool, models.SigninJwt, error) {
	if len(username) == 0 {
		return false, models.SigninJwt{}, customerrors.NewServiceError(http.StatusBadRequest, "username is required")
	}

	if len(password) == 0 {
		return false, models.SigninJwt{}, customerrors.NewServiceError(http.StatusBadRequest, "password is required")
	}

	return providers.IsValidUsernamePassword(ctx, username, password)
}

// Returns the signup
func Signup(ctx context.Context, username string, password string) (models.User, error) {
	if len(username) == 0 {
		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "username is required")
	}

	if len(password) == 0 {
		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "password is required")
	}

	return providers.Signup(ctx, username, password)
}
