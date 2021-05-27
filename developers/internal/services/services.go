package services

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/mikedelafuente/authful/developers/internal/config"
	"github.com/mikedelafuente/authful/developers/internal/models"
	"github.com/mikedelafuente/authful/developers/internal/repo"
	"github.com/mikedelafuente/authful/servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful/servertools/pkg/customerrors"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateDeveloper(ctx context.Context, username string, password string) (models.Developer, error) {
	if strings.TrimSpace(username) == "" {

		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	if strings.TrimSpace(password) == "" {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "password cannot be blank")
	}

	username = strings.ToLower(username)

	if !IsUniqueUsername(ctx, username) {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "username is not valid")
	}

	// bcrypt the password
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), config.GetConfig().Security.PasswordCostFactor)
	if err != nil {
		return models.Developer{}, err
	}

	return repo.CreateDeveloper(ctx, username, string(passwordBytes))
}

func GetDeveloperByUsername(ctx context.Context, username string) (models.Developer, error) {
	if strings.TrimSpace(username) == "" {
		return models.Developer{}, customerrors.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	return repo.GetDeveloperByUsername(ctx, username)
}

func GetDevelopers(ctx context.Context) ([]models.Developer, error) {
	return repo.GetDevelopers(ctx)
}

// Produces a JWT token for the user. Returns the token, the expiration time (UTC) and any error
func ProduceJwtTokenForDeveloper(ctx context.Context, username string, devId string) (string, time.Time, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().UTC().Add(30 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &customclaims.Claims{
		Username: username,
		SystemId: devId,
		Type:     "developer",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Id:        devId,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string

	tokenString, err := token.SignedString([]byte(config.GetConfig().Security.JwtKey))

	return tokenString, expirationTime, err
}

func IsUniqueUsername(ctx context.Context, username string) bool {
	user, err := repo.GetDeveloperByUsername(ctx, username)
	if err != nil {
		return false
	}

	// If the strings are equal, than the user exists...
	return !strings.EqualFold(user.Username, username)
}

func IsValidUsernamePassword(ctx context.Context, username string, password string) bool {
	user, bcryptPassword, err := repo.GetDeveloperWithPasswordByUsername(ctx, username)
	if err != nil {
		return false
	}

	if !strings.EqualFold(user.Username, username) {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(bcryptPassword), []byte(password))
	return err == nil
}
