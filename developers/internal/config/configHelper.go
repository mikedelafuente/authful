package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

var configOnce sync.Once
var configInstance *DeveloperServerConfig

var dbOnce sync.Once
var dbInstance *sql.DB

func GetConfig() *DeveloperServerConfig {
	configOnce.Do(func() {
		var err error
		if len(os.Getenv("WEB_SERVER_PORT")) == 0 {
			configInstance, err = getConfigInstanceFromFile()

		} else {
			configInstance, err = getConfigInstanceFromEnvironment()
		}
		if err != nil {
			panic(err)
		}
	})

	return configInstance
}

func getConfigInstanceFromEnvironment() (*DeveloperServerConfig, error) {
	log.Printf("Loading config from environment")

	var myConfig *DeveloperServerConfig = &DeveloperServerConfig{
		WebServer:      WebServerConfig{},
		DatabaseServer: DatabaseServerConfig{},
		Security:       SecurityConfig{},
	}

	// WEB SERVER
	myConfig.WebServer.Port = os.Getenv("WEB_SERVER_PORT")

	// SECURITY
	port, err := strconv.Atoi(os.Getenv("SECURITY_PASSWORD_COST_FACTOR"))
	if err != nil {
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

func getConfigInstanceFromFile() (*DeveloperServerConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	filePath := currDir + "/settings/config.json"
	log.Printf("Loading config from file: %s \n", filePath)
	// Load config from file system
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var myConfig *DeveloperServerConfig = &DeveloperServerConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		return nil, err

	}

	return myConfig, nil
}

func getDbConnectionInstance() (*sql.DB, error) {

	config := GetConfig()
	log.Printf("Instantiating database connection to :%s \n", config.DatabaseServer.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DatabaseServer.Username, config.DatabaseServer.Password, config.DatabaseServer.Host, config.DatabaseServer.Port, config.DatabaseServer.Database))

	return db, err
}

func GetDbConnection() *sql.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = getDbConnectionInstance()
		if err != nil {
			panic(err)
		}
	})

	return dbInstance
}
