package providers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/signin/internal/config"
	"github.com/mikedelafuente/authful/signin/internal/models"
)

func IsValidUsernamePassword(ctx context.Context, username string, password string) (bool, models.SigninJwt, error) {
	usersUri := config.GetConfig().Providers.UserServerUri
	if len(usersUri) == 0 {
		return false, models.SigninJwt{}, errors.New("user provider configured with empty url")
	}

	usersUri = usersUri + "/api/v1/account:signin"
	requestModel := models.SigninCredentials{
		Username: username,
		Password: password,
	}

	bodyBytes, statusCode, err := httptools.Post(ctx, usersUri, requestModel)
	if err != nil {
		if statusCode == http.StatusUnauthorized {
			return false, models.SigninJwt{}, customerrors.NewServiceError(statusCode, "authentication failed")
		} else {
			logger.Error(ctx, err)
			return false, models.SigninJwt{}, err
		}

	}

	var responseObject models.SigninJwt
	json.Unmarshal(bodyBytes, &responseObject)
	return true, responseObject, nil
}

func Signup(ctx context.Context, username string, password string) (models.User, error) {
	usersUri := config.GetConfig().Providers.UserServerUri
	if len(usersUri) == 0 {
		return models.User{}, errors.New("user provider configured with empty url")
	}

	usersUri = usersUri + "/api/v1/account:signup"
	requestModel := models.SigninCredentials{
		Username: username,
		Password: password,
	}

	bodyBytes, _, err := httptools.Post(ctx, usersUri, requestModel)
	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, err
	}

	var responseObject models.User
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject, nil
}
