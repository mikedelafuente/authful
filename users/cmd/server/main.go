package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weekendprojectapp/authful/serverutils"
	"github.com/weekendprojectapp/authful/users/internal/config"
	"github.com/weekendprojectapp/authful/users/internal/users/rest"

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
	openR.HandleFunc("/api/v1/account:signin", rest.AccountSigninPost).Methods(http.MethodPost)
	openR.HandleFunc("/api/v1/account:signup", rest.AccountSignupPost).Methods(http.MethodPost)

	// ------------ UNPROTECTED API ENDPOINTS ------------
	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/api/v1/users", rest.UsersGet).Methods(http.MethodGet)
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

func cookieJwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isValid := false

		cookie, err := r.Cookie("userSessionToken")
		if err != nil {
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

	var claims serverutils.Claims
	token, err := jwt.ParseWithClaims(rawToken, &claims, func(t *jwt.Token) (interface{}, error) {
		localClaim := t.Claims.(*serverutils.Claims)
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

	ctx := context.WithValue(r.Context(), serverutils.ContextKeySystemId, systemId)
	ctx = context.WithValue(ctx, serverutils.ContextKeySystemType, systemType)
	r = r.WithContext(ctx)

	return isValid, r
}
