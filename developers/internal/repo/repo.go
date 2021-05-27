package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/mikedelafuente/authful/developers/internal/config"
	"github.com/mikedelafuente/authful/developers/internal/models"

	"github.com/google/uuid"
)

// type User struct{}

// func New() *User {
// 	d := User{}
// 	return &d
// }

func CreateDeveloper(ctx context.Context, username string, password string) (models.Developer, error) {
	db := config.GetDbConnection()
	id := uuid.New().String()
	currentTime := time.Now().UTC()
	username = strings.ToLower(username)

	result, err := db.Exec("INSERT INTO developers (dev_id, username, password, create_datetime, update_datetime) VALUES (?, ?, ?, ?, ?)", id, username, password, currentTime, currentTime)
	if err != nil {
		return models.Developer{}, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return models.Developer{}, errors.New("failed to create a new user")
	}

	newDeveloper := models.Developer{
		Id:         id,
		Username:   username,
		CreateDate: currentTime,
		UpdateDate: currentTime,
	}

	return newDeveloper, nil
}
func GetDeveloperByUsername(ctx context.Context, username string) (models.Developer, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT dev_id, username, create_datetime, update_datetime FROM developers WHERE username = ? LIMIT 1", username)
	if err != nil {
		log.Print(err)
		return models.Developer{}, err
	}

	if result.Next() {
		return mapResultToDeveloper(result)
	} else {
		return models.Developer{}, nil
	}

}

func GetDeveloperWithPasswordByUsername(ctx context.Context, username string) (models.Developer, string, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT dev_id, username, create_datetime, update_datetime, password FROM developers WHERE username = ? LIMIT 1", username)
	if err != nil {
		log.Print(err)
		return models.Developer{}, "", err
	}

	if result.Next() {
		var user models.Developer = models.Developer{}
		var ntCreate sql.NullTime
		var ntUpdate sql.NullTime
		var password string

		err := result.Scan(&user.Id, &user.Username, &ntCreate, &ntUpdate, &password)

		if err != nil {
			log.Print(err)
			return models.Developer{}, "", err
		}

		if ntCreate.Valid {
			user.CreateDate = ntCreate.Time
		}

		if ntUpdate.Valid {
			user.UpdateDate = ntUpdate.Time
		}

		return user, password, err
	} else {
		return models.Developer{}, "", nil
	}

}

func GetDevelopers(ctx context.Context) ([]models.Developer, error) {
	devs := []models.Developer{}

	db := config.GetDbConnection()

	result, err := db.Query("SELECT dev_id, username, create_datetime, update_datetime FROM developers")

	if err != nil {
		return devs, err
	}

	for result.Next() {

		dev, err := mapResultToDeveloper(result)

		if err != nil {
			log.Print(err)
		} else {
			devs = append(devs, dev)
		}
	}
	return devs, nil
}

func mapResultToDeveloper(result *sql.Rows) (models.Developer, error) {

	var dev models.Developer = models.Developer{}
	var ntCreate sql.NullTime
	var ntUpdate sql.NullTime

	err := result.Scan(&dev.Id, &dev.Username, &ntCreate, &ntUpdate)

	if err != nil {
		log.Print(err)
		return models.Developer{}, err
	}

	if ntCreate.Valid {
		dev.CreateDate = ntCreate.Time
	}

	if ntUpdate.Valid {
		dev.UpdateDate = ntUpdate.Time
	}

	return dev, nil
}
