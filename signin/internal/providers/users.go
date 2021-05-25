package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/weekendprojectapp/authful/servertools"
	"github.com/weekendprojectapp/authful/signin/internal/config"
	"github.com/weekendprojectapp/authful/signin/pkg/models"
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

	requestBytes, err := servertools.MarshalFormat(requestModel)
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

	if !servertools.IsOkResponse(resp) {
		if resp.StatusCode == http.StatusUnauthorized {
			// Bad username and password
			return false, models.SigninJwt{}, nil
		}

		// TODO: try to extract out an error from the body
		return false, models.SigninJwt{}, servertools.NewServiceError(resp.StatusCode, "user service exception")
	}

	var responseObject models.SigninJwt
	json.Unmarshal(bodyBytes, &responseObject)

	return true, responseObject, nil
}
