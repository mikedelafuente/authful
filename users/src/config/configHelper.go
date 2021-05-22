package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var configOnce sync.Once
var configInstance *AuthfulConfig

var dbOnce sync.Once
var dbInstance *sql.DB

func GetAuthfulConfig() *AuthfulConfig {
	configOnce.Do(func() {
		var err error
		configInstance, err = getAuthfulConfigInstance()
		if err != nil {
			panic(err)
		}
	})

	return configInstance
}

func getAuthfulConfigInstance() (*AuthfulConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	fmt.Printf("Loading config from directory: %s \n", currDir)
	// Load config from file system
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var myConfig *AuthfulConfig = &AuthfulConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		return nil, err

	}

	if myConfig.WebServer.Address == "" {
		return nil, errors.New("empty web server address")
	}

	if myConfig.DatabaseServer.Address == "" {
		return nil, errors.New("empty database server address")
	}

	return myConfig, nil
}

func getDbConnectionInstance() (*sql.DB, error) {

	config := GetAuthfulConfig()
	fmt.Printf("Instantiating database connection to %s:%s \n", config.DatabaseServer.Address, config.DatabaseServer.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/authful_users?parseTime=true", config.DatabaseServer.Username, config.DatabaseServer.Password, config.DatabaseServer.Address, config.DatabaseServer.Port))

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
