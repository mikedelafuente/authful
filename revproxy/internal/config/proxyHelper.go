package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var (
	proxyOnce     sync.Once
	proxyInstance *ProxyConfig
)

func GetProxyConfig() *ProxyConfig {
	proxyOnce.Do(func() {
		var err error

		if len(os.Getenv("WEB_SERVER_PORT")) == 0 {
			proxyInstance, err = getProxyConfigInstanceFromFile()
		} else {
			tmpProxyInstance, fileErr := getProxyConfigInstanceFromFile()
			if fileErr != nil {
				fmt.Printf("ERROR: %s \n", fileErr)
				panic(err)
			}
			proxyInstance, err = getProxyConfigInstanceFromEnvironment(tmpProxyInstance)
		}
		if err != nil {
			fmt.Printf("ERROR: %s \n", err)
			panic(err)
		}
	})

	return proxyInstance
}
func getServiceUrlForName(serviceName string) string {
	switch serviceName {
	case "users":
		return os.Getenv("PROVIDERS_USER_SERVER_URI")
	case "developers":
		return os.Getenv("PROVIDERS_DEVELOPER_SERVER_URI")
	}

	return ""
}
func getProxyConfigInstanceFromEnvironment(tmpConfig *ProxyConfig) (*ProxyConfig, error) {
	fmt.Println("Loading developer server config from environment")

	for i := 0; i < len(tmpConfig.ProxyMaps); i++ {
		tmpConfig.ProxyMaps[i].ServiceBaseUrl = getServiceUrlForName(tmpConfig.ProxyMaps[i].Name)
	}

	return tmpConfig, nil
}

func getProxyConfigInstanceFromFile() (*ProxyConfig, error) {
	var err error

	currDir, _ := os.Getwd()
	filePath := currDir + "/settings/proxy.json"
	fmt.Printf("Loading config from file: %s \n", filePath)
	// Load config from file system
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		return nil, err
	}

	var myConfig *ProxyConfig = &ProxyConfig{}
	err = json.Unmarshal(f, &myConfig)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		return nil, err

	}
	return myConfig, nil
}
