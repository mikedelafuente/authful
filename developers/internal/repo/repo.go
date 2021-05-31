package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/developers/internal/config"
	"github.com/mikedelafuente/authful/developers/internal/models"

	"github.com/google/uuid"
)

func CreateDeveloper(ctx context.Context, userId string, organizationName string, contactEmail string, agreeToTermsOfService bool) (models.Developer, error) {
	db := config.GetDbConnection()
	devId := uuid.New().String()
	currentTime := time.Now().UTC()

	result, err := db.Exec("INSERT INTO developers (dev_id, user_id, organization_name, contact_email, agree_to_tos, create_datetime, update_datetime) VALUES (?, ?, ?, ?, ?, ?, ?)", devId, userId, organizationName, contactEmail, agreeToTermsOfService, currentTime, currentTime)
	if err != nil {
		logger.Error(ctx, err)
		return models.Developer{}, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return models.Developer{}, errors.New("failed to create a new developer")
	}

	newDeveloper := models.Developer{
		DeveloperId:      devId,
		UserId:           userId,
		OrganizationName: organizationName,
		ContactEmail:     contactEmail,
		CreateDate:       currentTime,
		UpdateDate:       currentTime,
	}

	return newDeveloper, nil
}
func GetDeveloperByUserId(ctx context.Context, userId string) (models.Developer, error) {
	db := config.GetDbConnection()
	result, err := db.Query("SELECT dev_id, user_id, organization_name, contact_email, create_datetime, update_datetime FROM developers WHERE user_id = ? LIMIT 1", userId)
	if err != nil {
		logger.Error(ctx, err)
		return models.Developer{}, err
	}

	if result.Next() {
		return mapResultToDeveloper(ctx, result)
	} else {
		return models.Developer{}, nil
	}
}

func GetDevelopers(ctx context.Context) ([]models.Developer, error) {
	devs := []models.Developer{}

	db := config.GetDbConnection()

	result, err := db.Query("SELECT dev_id, user_id, organization_name, contact_email, create_datetime, update_datetime FROM developers")

	if err != nil {
		logger.Error(ctx, err)
		return devs, err
	}

	for result.Next() {

		dev, err := mapResultToDeveloper(ctx, result)

		if err != nil {
			logger.Error(ctx, err)
		} else {
			devs = append(devs, dev)
		}
	}
	return devs, nil
}

func mapResultToDeveloper(ctx context.Context, result *sql.Rows) (models.Developer, error) {

	var dev models.Developer = models.Developer{}
	var ntCreate sql.NullTime
	var ntUpdate sql.NullTime

	err := result.Scan(&dev.DeveloperId, &dev.UserId, &dev.OrganizationName, &dev.ContactEmail, &ntCreate, &ntUpdate)

	if err != nil {
		logger.Error(ctx, err)
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
