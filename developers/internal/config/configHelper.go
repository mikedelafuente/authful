package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var configOnce sync.Once
var configInstance *DeveloperServerConfig

var dbOnce sync.Once
var dbInstance *sql.DB

func GetConfig() *DeveloperServerConfig {
	configOnce.Do(func() {
		var err error
		configInstance, err = getConfigInstanceFromFile()
		if err != nil {
			panic(err)
		}
	})

	return configInstance
}

func getConfigInstanceFromFile() (*DeveloperServerConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	log.Printf("Loading config from directory: %s \n", currDir)
	// Load config from file system
	f, err := ioutil.ReadFile("settings/config.json")
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
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DatabaseServer.Username, config.DatabaseServer.Password, "localhost", config.DatabaseServer.Port, config.DatabaseServer.DatabaseName))

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
