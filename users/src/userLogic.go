package main

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"weekendproject.app/authful/users/config"
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
		return userDto{}, errors.New("username cannot be blank")
	}

	username = strings.ToLower(username)

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
