package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	configOnce     sync.Once
	configInstance *UserServerConfig
)

var (
	dbOnce     sync.Once
	dbInstance *sql.DB
)

func GetConfig() *UserServerConfig {
	configOnce.Do(func() {
		var err error

		if len(os.Getenv("WEB_SERVER_PORT")) == 0 {
			configInstance, err = getConfigInstanceFromFile()
		} else {
			configInstance, err = getConfigInstanceFromEnvironment()
		}
		if err != nil {
			fmt.Printf("ERROR: %s \n", err)
			panic(err)
		}
	})

	return configInstance
}

func getConfigInstanceFromEnvironment() (*UserServerConfig, error) {
	fmt.Println("Loading user server config from environment")

	var myConfig *UserServerConfig = &UserServerConfig{
		WebServer:      WebServerConfig{},
		DatabaseServer: DatabaseServerConfig{},
		Security:       SecurityConfig{},
	}
	myConfig.LogLevel = os.Getenv("AUTHFUL_LOG_LEVEL")

	// WEB SERVER
	myConfig.WebServer.Port = os.Getenv("WEB_SERVER_PORT")
	myConfig.WebServer.CORSOriginAllowed = parseCommaDelimitedStringToArray(os.Getenv("CORS_ORIGIN_ALLOWED"))
	myConfig.WebServer.CORSAllowedHeaders = parseCommaDelimitedStringToArray(os.Getenv("CORS_ALLOWED_HEADERS"))
	myConfig.WebServer.CORSAllowedMethods = parseCommaDelimitedStringToArray(os.Getenv("CORS_ALLOWED_METHODS"))

	// SECURITY
	port, err := strconv.Atoi(os.Getenv("SECURITY_PASSWORD_COST_FACTOR"))
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
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

func parseCommaDelimitedStringToArray(v string) []string {
	if len(v) == 0 {
		return []string{}
	}

	result := []string{}
	parts := strings.Split(v, ",")
	for _, item := range parts {
		result = append(result, strings.TrimSpace(item))
	}

	return result
}

func getConfigInstanceFromFile() (*UserServerConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	filePath := currDir + "/settings/config.json"
	fmt.Printf("Loading config from file: %s \n", filePath)
	// Load config from file system
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		return nil, err
	}

	var myConfig *UserServerConfig = &UserServerConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		return nil, err
	}
	os.Setenv("AUTHFUL_LOG_LEVEL", myConfig.LogLevel)

	return myConfig, nil
}

func getDbConnectionInstance() (*sql.DB, error) {
	config := GetConfig()
	fmt.Printf("Instantiating database connection to :%v \n", config.DatabaseServer.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DatabaseServer.Username, config.DatabaseServer.Password, config.DatabaseServer.Host, config.DatabaseServer.Port, config.DatabaseServer.Database))

	return db, err
}

func GetDbConnection() *sql.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = getDbConnectionInstance()
		if err != nil {
			fmt.Printf("ERROR: %s \n", err)
			panic(err)
		}
	})

	return dbInstance
}
