package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var configOnce sync.Once
var configInstance *ServerConfig

func GetConfig() *ServerConfig {
	configOnce.Do(func() {
		var err error
		configInstance, err = getConfigInstanceFromFile()
		if err != nil {
			panic(err)
		}
	})

	return configInstance
}

func getConfigInstanceFromFile() (*ServerConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	log.Printf("Loading config from directory: %s \n", currDir)
	// Load config from file system
	f, err := ioutil.ReadFile(currDir + "/settings/config.json")
	if err != nil {
		return nil, err
	}

	var myConfig *ServerConfig = &ServerConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		return nil, err

	}

	if myConfig.WebServer.Address == "" {
		return nil, errors.New("empty web server address")
	}

	if myConfig.Providers.UserServerUri == "" {
		return nil, errors.New("empty user service uri")
	}

	return myConfig, nil
}
