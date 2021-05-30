package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
)

var configOnce sync.Once
var configInstance *UserServerConfig

var dbOnce sync.Once
var dbInstance *sql.DB

func GetConfig() *UserServerConfig {
	configOnce.Do(func() {
		var err error
		if len(os.Getenv("WEB_SERVER_PORT")) == 0 {
			configInstance, err = getConfigInstanceFromFile()

		} else {
			configInstance, err = getConfigInstanceFromEnvironment()
		}
		if err != nil {
			logger.Error(err)
			panic(err)
		}
	})

	return configInstance
}

func getConfigInstanceFromEnvironment() (*UserServerConfig, error) {
	logger.Printf("Loading config from environment")

	var myConfig *UserServerConfig = &UserServerConfig{
		WebServer:      WebServerConfig{},
		DatabaseServer: DatabaseServerConfig{},
		Security:       SecurityConfig{},
	}

	myConfig.IsDebug, _ = strconv.ParseBool(os.Getenv("IS_DEBUG"))

	// WEB SERVER
	myConfig.WebServer.Port = os.Getenv("WEB_SERVER_PORT")

	// SECURITY
	port, err := strconv.Atoi(os.Getenv("SECURITY_PASSWORD_COST_FACTOR"))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	myConfig.Security.PasswordCostFactor = port
	myConfig.Security.JwtKey = os.Getenv("SECURITY_JWT_KEY")

	// DATABASE SERVER
	myConfig.DatabaseServer.Host = os.Getenv("DATABASE_SERVER_HOST")
	myConfig.DatabaseServer.Database = os.Getenv("DATABASE_SERVER_DATABASE")
	myConfig.DatabaseServer.Password = os.Getenv("DATABASE_SERVER_PASSWORD")
	myConfig.DatabaseServer.Port = os.Getenv("DATABASE_SERVER_PORT")
	myConfig.DatabaseServer.Username = os.Getenv("DATABASE_SERVER_USERNAME")

	return myConfig, nil
}

func getConfigInstanceFromFile() (*UserServerConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	filePath := currDir + "/settings/config.json"
	logger.Printf("Loading config from file: %s \n", filePath)
	// Load config from file system
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var myConfig *UserServerConfig = &UserServerConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	os.Setenv("IS_DEBUG", fmt.Sprintf("%t", myConfig.IsDebug))
	return myConfig, nil
}

func getDbConnectionInstance() (*sql.DB, error) {

	config := GetConfig()
	logger.Printf("Instantiating database connection to :%s \n", config.DatabaseServer.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DatabaseServer.Username, config.DatabaseServer.Password, config.DatabaseServer.Host, config.DatabaseServer.Port, config.DatabaseServer.Database))

	return db, err
}

func GetDbConnection() *sql.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = getDbConnectionInstance()
		if err != nil {
			logger.Error(err)
			panic(err)
		}
	})

	return dbInstance
}
