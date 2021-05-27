package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikedelafuente/authful/developers/internal/config"
	"github.com/mikedelafuente/authful/developers/internal/controllers"
	"github.com/mikedelafuente/authful/servertools/pkg/customclaims"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	startTime = time.Now()
	fmt.Printf("Process started at %s\n", startTime)
	config.GetConfig()       // just attempt to get the config at startup
	config.GetDbConnection() // just attempt to connect to the database at startup
}

func main() {
	myConfig := config.GetConfig()
	fmt.Printf("\n\nAuthful: User Server running at %s:%v\n\n", myConfig.WebServer.Address, myConfig.WebServer.Port)
	setupRequestHandlers()
}

func setupRequestHandlers() {

	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost).Subrouter()
	openR.HandleFunc("/api/v1/account:signin", controllers.DeveloperSigninPost).Methods(http.MethodPost)
	openR.HandleFunc("/api/v1/account:signup", controllers.DeveloperSignupPost).Methods(http.MethodPost)

	// ------------ UNPROTECTED API ENDPOINTS ------------
	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/api/v1/users", controllers.DevelopersGet).Methods(http.MethodGet)
	secureUserR.Use(bearerJwtHandler)

	myConfig := config.GetConfig()

	defer dbShutdown()
	err := http.ListenAndServe(fmt.Sprintf("%s:%v", myConfig.WebServer.Address, myConfig.WebServer.Port), myRouter)
	endTime := time.Now()
	fmt.Printf("Process stopped at %s\n", endTime)
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Server uptime was: %s", elapsed)
	log.Fatal(err)
}

func dbShutdown() {
	fmt.Println("shutting down database")
	db := config.GetDbConnection()
	db.Close()
}

func bearerJwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// uh.logger.Debug("validating access token")
		authHeader := r.Header.Values("Authorization")
		isValid := false

		if len(authHeader) > 0 {
			if strings.HasPrefix(authHeader[0], "Bearer ") {
				parts := strings.Split(authHeader[0], " ")
				rawToken := parts[1]

				isValid, r = processToken(rawToken, r)
			}
		}

		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func processToken(rawToken string, r *http.Request) (bool, *http.Request) {
	systemId := ""
	systemType := ""
	isValid := false

	var claims customclaims.Claims
	token, err := jwt.ParseWithClaims(rawToken, &claims, func(t *jwt.Token) (interface{}, error) {
		localClaim := t.Claims.(*customclaims.Claims)
		systemId = localClaim.SystemId
		systemType = localClaim.Type
		return []byte(config.GetConfig().Security.JwtKey), nil
	})

	if err == nil {
		if token.Valid {
			isValid = true
		}
	} else {
		fmt.Println("Error happened: " + err.Error())
	}

	ctx := context.WithValue(r.Context(), customclaims.ContextKeySystemId, systemId)
	ctx = context.WithValue(ctx, customclaims.ContextKeySystemType, systemType)
	r = r.WithContext(ctx)

	return isValid, r
}