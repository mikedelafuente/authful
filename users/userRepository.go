package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"weekendprojectapp/authful/users/config"

	"github.com/google/uuid"
)

type userRepository struct{}

func newUserRepository() *userRepository {
	d := userRepository{}
	return &d
}

func (d *userRepository) createUser(ctx context.Context, username string, password string) (userDto, error) {
	db := config.GetDbConnection()
	id := uuid.New().String()
	currentTime := time.Now().UTC()
	username = strings.ToLower(username)

	result, err := db.Exec("INSERT INTO users (id, username, password, create_datetime, update_datetime) VALUES (?, ?, ?, ?, ?)", id, username, password, currentTime, currentTime)
	if err != nil {
		return userDto{}, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return userDto{}, errors.New("failed to create a new user")
	}

	newUser := userDto{
		Id:         id,
		Username:   username,
		CreateDate: currentTime,
		UpdateDate: currentTime,
	}

	return newUser, nil
}
func (d *userRepository) getUserByUsername(ctx context.Context, username string) (userDto, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT id, username, create_datetime, update_datetime FROM users WHERE username = ? LIMIT 1", username)
	if err != nil {
		log.Print(err)
		return userDto{}, err
	}

	if result.Next() {
		return mapResultToUser(result)
	} else {
		return userDto{}, nil
	}

}

func (d *userRepository) getUsers(ctx context.Context) ([]userDto, error) {
	users := []userDto{}

	db := config.GetDbConnection()

	result, err := db.Query("SELECT id, username, create_datetime, update_datetime FROM users")

	if err != nil {
		return users, err
	}

	for result.Next() {

		user, err := mapResultToUser(result)

		if err != nil {
			log.Print(err)
		} else {
			users = append(users, user)
		}
	}
	return users, nil
}

func mapResultToUser(result *sql.Rows) (userDto, error) {

	var user userDto = userDto{}
	var ntCreate sql.NullTime
	var ntUpdate sql.NullTime

	err := result.Scan(&user.Id, &user.Username, &ntCreate, &ntUpdate)

	if err != nil {
		log.Print(err)
		return userDto{}, err
	}

	if ntCreate.Valid {
		user.CreateDate = ntCreate.Time
	}

	if ntUpdate.Valid {
		user.UpdateDate = ntUpdate.Time
	}

	return user, nil
}
