package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful/signin/internal/config"
	"github.com/mikedelafuente/authful/signin/internal/controllers"
	"github.com/mikedelafuente/authful/signin/internal/logger"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	log.SetOutput(os.Stdout)
	startTime = time.Now()
	logger.Printf("Process started at %s\n", startTime)
	config.GetConfig() // just attempt to get the config at startup
}

func main() {
	myConfig := config.GetConfig()
	logger.Printf("\n\nAuthful: Signin Server running at :%v\n\n", myConfig.WebServer.Port)
	logger.Printf("DEBUG MODE: %t", config.GetConfig().IsDebug)
	setupRequestHandlers()
}

func setupRequestHandlers() {
	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	openR.HandleFunc("/login", controllers.DisplayLogin).Methods(http.MethodGet)
	openR.HandleFunc("/login", controllers.ProcessLogin).Methods(http.MethodPost)
	openR.HandleFunc("/signup", controllers.DisplaySignup).Methods(http.MethodGet)
	openR.HandleFunc("/signup", controllers.ProcessSignup).Methods(http.MethodPost)

	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/", controllers.Index).Methods(http.MethodGet)
	secureUserR.HandleFunc("/profile", controllers.GetProfile).Methods(http.MethodGet)
	secureUserR.Use(cookieJwtHandler)

	fileR := myRouter.Methods(http.MethodGet).Subrouter()
	fileServer := http.FileServer(http.Dir("./Static"))
	fileR.PathPrefix("/").Handler(http.StripPrefix("/resources", fileServer))

	// openR.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	// openR.HandleFunc("/", renderTemplate)

	myConfig := config.GetConfig()
	err := http.ListenAndServe(fmt.Sprintf(":%v", myConfig.WebServer.Port), myRouter)

	endTime := time.Now()
	logger.Printf("Process stopped at %s\n", endTime)
	elapsed := endTime.Sub(startTime)
	logger.Printf("Server uptime was: %s", elapsed)
	logger.Fatal(err)

}

func cookieJwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isValid := true

		cookie, err := r.Cookie("userSessionToken")
		if err != nil {
			logger.Println(err)
			// Redirect
			isValid = false
		}

		if isValid {
			rawToken := cookie.Value
			isValid, r = processToken(rawToken, r)
		}

		if !isValid {
			var loginRedirectUri = url.QueryEscape(r.URL.String())
			http.Redirect(w, r, "/login?redirect_uri="+loginRedirectUri, http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func processToken(rawToken string, r *http.Request) (bool, *http.Request) {
	userId := ""
	isValid := false

	var claims customclaims.Claims
	token, err := jwt.ParseWithClaims(rawToken, &claims, func(t *jwt.Token) (interface{}, error) {
		localClaim := t.Claims.(*customclaims.Claims)
		userId = localClaim.UserId
		return []byte(config.GetConfig().Security.JwtKey), nil
	})

	if err == nil {
		if token.Valid {
			isValid = true
		}
	} else {
		logger.Println("Error happened: " + err.Error())
	}

	ctx := context.WithValue(r.Context(), customclaims.ContextKeyUserId, userId)
	r = r.WithContext(ctx)

	return isValid, r
}
