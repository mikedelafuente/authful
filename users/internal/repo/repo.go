package repo

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/users/internal/config"
	"github.com/mikedelafuente/authful/users/internal/models"

	"github.com/google/uuid"
)

// type User struct{}

// func New() *User {
// 	d := User{}
// 	return &d
// }

func CreateUser(ctx context.Context, username string, password string) (models.User, error) {
	db := config.GetDbConnection()
	id := uuid.New().String()
	currentTime := time.Now().UTC()
	username = strings.ToLower(username)

	result, err := db.Exec("INSERT INTO users (user_id, username, password, create_datetime, update_datetime) VALUES (?, ?, ?, ?, ?)", id, username, password, currentTime, currentTime)
	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return models.User{}, errors.New("failed to create a new user")
	}

	newUser := models.User{
		UserId:     id,
		Username:   username,
		CreateDate: currentTime,
		UpdateDate: currentTime,
	}

	return newUser, nil
}
func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT user_id, username, create_datetime, update_datetime FROM users WHERE username = ? LIMIT 1", username)
	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, err
	}

	if result.Next() {
		return mapResultToUser(ctx, result)
	} else {
		return models.User{}, nil
	}

}

func GetUserWithPasswordByUsername(ctx context.Context, username string) (models.User, string, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT user_id, username, create_datetime, update_datetime, password FROM users WHERE username = ? LIMIT 1", username)
	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, "", err
	}

	if result.Next() {
		var user models.User = models.User{}
		var ntCreate sql.NullTime
		var ntUpdate sql.NullTime
		var password string

		err := result.Scan(&user.UserId, &user.Username, &ntCreate, &ntUpdate, &password)

		if err != nil {
			logger.Error(ctx, err)
			return models.User{}, "", err
		}

		if ntCreate.Valid {
			user.CreateDate = ntCreate.Time
		}

		if ntUpdate.Valid {
			user.UpdateDate = ntUpdate.Time
		}

		return user, password, err
	} else {
		return models.User{}, "", nil
	}

}

func GetUsers(ctx context.Context) ([]models.User, error) {
	users := []models.User{}

	db := config.GetDbConnection()

	result, err := db.Query("SELECT user_id, username, create_datetime, update_datetime FROM users")

	if err != nil {
		logger.Error(ctx, err)
		return users, err
	}

	for result.Next() {

		user, err := mapResultToUser(ctx, result)

		if err != nil {
			logger.Error(ctx, err)
		} else {
			users = append(users, user)
		}
	}
	return users, nil
}

func mapResultToUser(ctx context.Context, result *sql.Rows) (models.User, error) {

	var user models.User = models.User{}
	var ntCreate sql.NullTime
	var ntUpdate sql.NullTime

	err := result.Scan(&user.UserId, &user.Username, &ntCreate, &ntUpdate)

	if err != nil {
		logger.Error(ctx, err)
		return models.User{}, err
	}

	if ntCreate.Valid {
		user.CreateDate = ntCreate.Time
	}

	if ntUpdate.Valid {
		user.UpdateDate = ntUpdate.Time
	}

	return user, nil
}
