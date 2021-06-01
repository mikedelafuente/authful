package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var (
	configOnce     sync.Once
	configInstance *ServerConfig
)

func GetConfig() *ServerConfig {
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

func getConfigInstanceFromEnvironment() (*ServerConfig, error) {
	fmt.Println("Loading signin server config from environment")

	var myConfig *ServerConfig = &ServerConfig{
		WebServer: WebServerConfig{},
		Providers: ProvidersConfig{},
		Security:  SecurityConfig{},
	}
	myConfig.LogLevel = os.Getenv("AUTHFUL_LOG_LEVEL")

	// WEB SERVER
	myConfig.WebServer.Port = os.Getenv("WEB_SERVER_PORT")
	myConfig.WebServer.CORSOriginAllowed = parseCommaDelimitedStringToArray(os.Getenv("CORS_ORIGIN_ALLOWED"))
	myConfig.WebServer.CORSAllowedHeaders = parseCommaDelimitedStringToArray(os.Getenv("CORS_ALLOWED_HEADERS"))
	myConfig.WebServer.CORSAllowedMethods = parseCommaDelimitedStringToArray(os.Getenv("CORS_ALLOWED_METHODS"))

	// SECURITY
	myConfig.Security.JwtKey = os.Getenv("SECURITY_JWT_KEY")

	// DATABASE SERVER
	myConfig.Providers.DeveloperServerUri = os.Getenv("PROVIDERS_DEVELOPER_SERVER_URI")
	myConfig.Providers.UserServerUri = os.Getenv("PROVIDERS_USER_SERVER_URI")

	/*
	  CORS_ORIGIN_ALLOWED: "http://localhost:8080, http://localhost:3000"
	  CORS_ALLOWED_HEADERS: "Accept,Cache,Content-Type,Authorize,Access-Control-Allow-Origin,X-Requested-With,X-Auth-Token,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,x-trace-id"
	  CORS_ALLOWED_METHODS: "GET,HEAD,POST,PUT,PATCH,OPTIONS"
	*/
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

func getConfigInstanceFromFile() (*ServerConfig, error) {
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

	var myConfig *ServerConfig = &ServerConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		return nil, err
	}
	os.Setenv("AUTHFUL_LOG_LEVEL", myConfig.LogLevel)

	if myConfig.Providers.UserServerUri == "" {
		return nil, errors.New("empty user service uri")
	}

	return myConfig, nil
}
