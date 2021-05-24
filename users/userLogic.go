package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/weekendprojectapp/authful/serverutils"
	"github.com/weekendprojectapp/authful/users/config"

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

func (l *userLogic) getUsers(ctx context.Context) ([]userDto, error) {
	return l.repo.getUsers(ctx)
}

func (l *userLogic) isUniqueUsername(ctx context.Context, username string) bool {
	user, err := l.repo.getUserByUsername(ctx, username)
	if err != nil {
		return false
	}

	// If the strings are equal, than the user exists...
	return !strings.EqualFold(user.Username, username)
}
