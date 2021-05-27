package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mikedelafuente/authful/servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful/servertools/pkg/httptools"
	"github.com/mikedelafuente/authful/signin/internal/config"
	"github.com/mikedelafuente/authful/signin/internal/models"
)

func IsValidUsernamePassword(ctx context.Context, username string, password string) (bool, models.SigninJwt, error) {
	usersUri := config.GetConfig().Providers.UserServerUri
	if len(usersUri) == 0 {
		return false, models.SigninJwt{}, errors.New("user provider configured with empty url")
	}

	usersUri = usersUri + "/api/v1/account:signin"
	client := &http.Client{}
	requestModel := models.SigninCredentials{
		Username: username,
		Password: password,
	}

	requestBytes, err := httptools.MarshalFormat(requestModel)
	if err != nil {
		return false, models.SigninJwt{}, err
	}

	req, err := http.NewRequest("POST", usersUri, bytes.NewBuffer(requestBytes))
	if err != nil {
		return false, models.SigninJwt{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, models.SigninJwt{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, models.SigninJwt{}, err
	}

	if !httptools.IsOkResponse(resp) {

		errorMessage := httptools.ExtractErrorMessageFromJsonBytes(bodyBytes, "authentication failed")

		if resp.StatusCode == http.StatusUnauthorized {
			return false, models.SigninJwt{}, customerrors.NewServiceError(http.StatusUnauthorized, errorMessage)
		}

		return false, models.SigninJwt{}, customerrors.NewServiceError(resp.StatusCode, errorMessage)
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
	client := &http.Client{}
	requestModel := models.SigninCredentials{
		Username: username,
		Password: password,
	}

	requestBytes, err := httptools.MarshalFormat(requestModel)
	if err != nil {
		return models.User{}, err
	}

	req, err := http.NewRequest("POST", usersUri, bytes.NewBuffer(requestBytes))
	if err != nil {
		return models.User{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, err
	}

	if !httptools.IsOkResponse(resp) {
		// TODO: try to extract out an error from the body
		errorMessage := httptools.ExtractErrorMessageFromJsonBytes(bodyBytes, "user service exception")

		return models.User{}, customerrors.NewServiceError(resp.StatusCode, errorMessage)
	}

	var responseObject models.User
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject, nil
}
