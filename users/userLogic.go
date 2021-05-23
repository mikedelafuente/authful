package main

import (
	"context"
	"errors"
	"strings"

	"weekendprojectapp/authful/users/config"

	"weekendprojectapp/serverutilities"

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
		return userDto{}, errors.New("username cannot be blank")
	}

	if strings.TrimSpace(password) == "" {
		return userDto{}, serverutilities.NewServiceError(300, "test")
	}

	username = strings.ToLower(username)

	if !l.isUniqueUsername(ctx, username) {
		return userDto{}, errors.New("username is not valid")
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
