package services

import (
	"context"
	"net/http"
	"strings"

	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/developers/internal/models"
	"github.com/mikedelafuente/authful/developers/internal/repo"
)

func CreateDeveloper(ctx context.Context, userId string, organizationName string, contactEmail string, agreeToTermsOfService bool) (models.Developer, error) {
	if strings.TrimSpace(userId) == "" {

		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "user_id cannot be blank")
	}

	if strings.TrimSpace(organizationName) == "" {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "password cannot be blank")
	}

	if strings.TrimSpace(contactEmail) == "" {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "contact email cannot be blank")
	}

	if !agreeToTermsOfService {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "developer accounts require that you agree to the terms of service")
	}

	if !IsUniqueUserId(ctx, userId) {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "developer account already exists for the user")
	}

	return repo.CreateDeveloper(ctx, userId, organizationName, contactEmail, agreeToTermsOfService)
}

func GetDevelopers(ctx context.Context) ([]models.Developer, error) {
	return repo.GetDevelopers(ctx)
}

func IsUniqueUserId(ctx context.Context, userId string) bool {
	user, err := repo.GetDeveloperByUserId(ctx, userId)
	if err != nil {
		logger.Error(err)
		return false
	}

	// If the strings are equal, than the user exists...
	return !strings.EqualFold(user.UserId, userId)
}
