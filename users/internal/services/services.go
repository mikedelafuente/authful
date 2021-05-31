package services

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful-servertools/pkg/customerrors"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/users/internal/config"
	"github.com/mikedelafuente/authful/users/internal/models"
	"github.com/mikedelafuente/authful/users/internal/repo"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, username string, password string) (models.User, error) {
	if strings.TrimSpace(username) == "" {

		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	if strings.TrimSpace(password) == "" {
		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "password cannot be blank")
	}

	username = strings.ToLower(username)

	if !IsUniqueUsername(ctx, username) {
		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "username is not valid")
	}

	// bcrypt the password
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), config.GetConfig().Security.PasswordCostFactor)
	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, err
	}

	return repo.CreateUser(ctx, username, string(passwordBytes))
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	if strings.TrimSpace(username) == "" {
		return models.User{}, customerrors.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	return repo.GetUserByUsername(ctx, username)
}

func GetUsers(ctx context.Context) ([]models.User, error) {
	return repo.GetUsers(ctx)
}

// Produces a JWT token for the user. Returns the token, the expiration time (UTC) and any error
func ProduceJwtTokenForUser(ctx context.Context, username string, userId string) (string, time.Time, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().UTC().Add(30 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &customclaims.Claims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Id:        userId,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string

	tokenString, err := token.SignedString([]byte(config.GetConfig().Security.JwtKey))

	return tokenString, expirationTime, err
}

func IsUniqueUsername(ctx context.Context, username string) bool {
	user, err := repo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	// If the strings are equal, than the user exists...
	return !strings.EqualFold(user.Username, username)
}

func IsValidUsernamePassword(ctx context.Context, username string, password string) bool {
	user, bcryptPassword, err := repo.GetUserWithPasswordByUsername(ctx, username)
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	if !strings.EqualFold(user.Username, username) {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(bcryptPassword), []byte(password))
	return err == nil
}
