package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful/users/internal/config"
	"github.com/mikedelafuente/authful/users/internal/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	startTime = time.Now()
	log.Printf("Process started at %s\n", startTime)
	config.GetConfig()       // just attempt to get the config at startup
	config.GetDbConnection() // just attempt to connect to the database at startup
}

func main() {
	myConfig := config.GetConfig()
	log.Printf("\n\nAuthful: User Server running at %s:%v\n\n", myConfig.WebServer.Address, myConfig.WebServer.Port)
	setupRequestHandlers()
}

func setupRequestHandlers() {

	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost).Subrouter()
	openR.HandleFunc("/api/v1/account:signin", controllers.AccountSigninPost).Methods(http.MethodPost)
	openR.HandleFunc("/api/v1/account:signup", controllers.AccountSignupPost).Methods(http.MethodPost)

	// ------------ UNPROTECTED API ENDPOINTS ------------
	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/api/v1/users", controllers.UsersGet).Methods(http.MethodGet)
	secureUserR.Use(bearerJwtHandler)

	myConfig := config.GetConfig()

	defer dbShutdown()
	err := http.ListenAndServe(fmt.Sprintf("%s:%v", myConfig.WebServer.Address, myConfig.WebServer.Port), myRouter)
	endTime := time.Now()
	log.Printf("Process stopped at %s\n", endTime)
	elapsed := endTime.Sub(startTime)
	log.Printf("Server uptime was: %s", elapsed)
	log.Fatal(err)
}

func dbShutdown() {
	log.Println("shutting down database")
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
		log.Println("Error happened: " + err.Error())
	}

	ctx := context.WithValue(r.Context(), customclaims.ContextKeyUserId, userId)
	r = r.WithContext(ctx)

	return isValid, r
}
