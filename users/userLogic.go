package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/weekendprojectapp/authful/serverutils"
	"github.com/weekendprojectapp/authful/users/config"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userLogic struct {
	repo userRepository
}

func newUserLogic() *userLogic {
	d := userLogic{
		repo: *newUserRepository(),
	}
	return &d
}

func (l *userLogic) createUser(ctx context.Context, username string, password string) (userDto, error) {
	if strings.TrimSpace(username) == "" {
		return userDto{}, serverutils.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	if strings.TrimSpace(password) == "" {
		return userDto{}, serverutils.NewServiceError(http.StatusBadRequest, "password cannot be blank")
	}

	username = strings.ToLower(username)

	if !l.isUniqueUsername(ctx, username) {
		return userDto{}, serverutils.NewServiceError(http.StatusBadRequest, "username is not valid")
	}

	// bcrypt the password
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), config.GetAuthfulConfig().Security.PasswordCostFactor)
	if err != nil {
		return userDto{}, err
	}

	return l.repo.createUser(ctx, username, string(passwordBytes))
}

func (l *userLogic) getUserByUsername(ctx context.Context, username string) (userDto, error) {
	if strings.TrimSpace(username) == "" {
		return userDto{}, serverutils.NewServiceError(http.StatusBadRequest, "username cannot be blank")
	}

	return l.repo.getUserByUsername(ctx, username)
}

func (l *userLogic) getUsers(ctx context.Context) ([]userDto, error) {
	return l.repo.getUsers(ctx)
}

// Produces a JWT token for the user. Returns the token, the expiration time (UTC) and any error
func (l *userLogic) produceJwtTokenForUser(ctx context.Context, username string, userId string) (string, time.Time, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().UTC().Add(30 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &serverutils.Claims{
		Username: username,
		SystemId: userId,
		Type:     "user",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Id:        userId,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string

	tokenString, err := token.SignedString([]byte(config.GetAuthfulConfig().Security.JwtKey))

	return tokenString, expirationTime, err
}

func (l *userLogic) isUniqueUsername(ctx context.Context, username string) bool {
	user, err := l.repo.getUserByUsername(ctx, username)
	if err != nil {
		return false
	}

	// If the strings are equal, than the user exists...
	return !strings.EqualFold(user.Username, username)
}

func (l *userLogic) isValidUsernamePassword(ctx context.Context, username string, password string) bool {
	user, bcryptPassword, err := l.repo.getUserWithPasswordByUsername(ctx, username)
	if err != nil {
		return false
	}

	if !strings.EqualFold(user.Username, username) {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(bcryptPassword), []byte(password))
	return err == nil
}
