package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/revproxy/internal/config"
)

type requestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

type ProxyInfo struct {
	ServiceBaseUrl string
	Path           string
	IsSecure       bool
}

var serviceMap map[string]ProxyInfo

func init() {
	log.SetOutput(os.Stdout)
	config.GetConfig()
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	parsedUrl, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(parsedUrl)

	// Update the headers to allow for SSL redirection
	req.URL.Host = parsedUrl.Host
	req.URL.Scheme = parsedUrl.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = parsedUrl.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	logger.Verbose(req.Context(), fmt.Sprintf("Sending request to server. Host: %v | Scheme: %v | %v ", parsedUrl.Host, parsedUrl.Scheme, parsedUrl))
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	logger.Verbose(r.Context(), fmt.Sprintf("Request recieved: %s %s", r.Method, r.URL))
	payload := parseRequest(r)
	logger.Verbose(r.Context(), "Proxy condition: "+payload.ProxyCondition)
	proxyInfo := getProxyInfo(r.Context(), payload.ProxyCondition)

	logger.Verbose(r.Context(), "Destination url: "+proxyInfo.ServiceBaseUrl)

	if proxyInfo.IsSecure {
		logger.Verbose(r.Context(), "Is secure endpoint")
	} else {
		logger.Verbose(r.Context(), "Is unsecure endpoint")
	}

	serveReverseProxy(proxyInfo.ServiceBaseUrl, w, r)
}

func parseRequest(r *http.Request) requestPayloadStruct {
	rp := requestPayloadStruct{}

	path := strings.ToUpper(strings.TrimSpace(r.URL.Path))
	if path[0] == byte('/') {
		parts := strings.Split(path, "/")
		if len(parts) < 4 {
			logger.Warn(r.Context(), "Invalid path. Not enough parts.")
			return rp
		}
		if strings.Contains(parts[3], ":") {
			parts[3] = strings.Split(parts[3], ":")[0]
		}
		rp.ProxyCondition = fmt.Sprintf("/%s/%s/%s", parts[1], parts[2], parts[3])
	} else {
		logger.Warn(r.Context(), "Invalid path. Does not start with a '/'")
	}

	return rp
}

func initializeServiceMap() {
	fmt.Printf("\n\nInitializing service map\n\n")
	serviceMap = make(map[string]ProxyInfo)
	for _, proxyMap := range config.GetConfig().ProxyMaps {
		for _, path := range proxyMap.Paths {
			info := ProxyInfo{
				ServiceBaseUrl: proxyMap.ServiceBaseUrl,
				Path:           strings.ToUpper(path.Path),
				IsSecure:       path.IsSecure,
			}
			serviceMap[strings.ToUpper(path.Path)] = info
		}
	}
}

func getProxyInfo(ctx context.Context, proxyConditionRaw string) ProxyInfo {
	proxyCondition := strings.ToUpper(proxyConditionRaw)

	proxyInfo := serviceMap[proxyCondition]
	if len(proxyInfo.ServiceBaseUrl) == 0 {
		logger.Warn(ctx, fmt.Sprintf("No Service URL found for [%v]", proxyCondition))
	} else {
		logger.Verbose(ctx, fmt.Sprintf("Service URL found for [%v] : %v", proxyCondition, proxyInfo))
	}

	return proxyInfo
}

func main() {
	fmt.Printf("\n\nAuthful: Reverse Proxy Server\n\n")
	fmt.Printf("Log Level: %s\n", logger.GetLogLevel())

	initializeServiceMap()

	fmt.Printf("\n\nAuthful: Reverse proxy Server running at %s:%v\n\n", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port)

	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port), nil); err != nil {
		panic(err)
	}
}
