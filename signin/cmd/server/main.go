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
	"github.com/weekendprojectapp/authful/users/config"
	"github.com/weekendprojectapp/authful/users/internal/web"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	startTime = time.Now()
	fmt.Printf("Process started at %s\n", startTime)
	config.GetAuthfulConfig() // just attempt to get the config at startup
}

func main() {
	myConfig := config.GetAuthfulConfig()
	fmt.Printf("\n\nAuthful: User Server running at %s:%v\n\n", myConfig.WebServer.Address, myConfig.WebServer.Port)
	setupRequestHandlers()
}

func setupRequestHandlers() {

	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost).Subrouter()
	openR.HandleFunc("/login", web.DisplayLogin).Methods(http.MethodGet)
	openR.HandleFunc("/login", web.AuthorizeUser).Methods(http.MethodPost)

	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/profile", web.GetProfile).Methods(http.MethodGet)
	secureUserR.Use(cookieJwtHandler)

	myConfig := config.GetAuthfulConfig()

	err := http.ListenAndServe(fmt.Sprintf("%s:%v", myConfig.WebServer.Address, myConfig.WebServer.Port), myRouter)
	endTime := time.Now()
	fmt.Printf("Process stopped at %s\n", endTime)
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Server uptime was: %s", elapsed)
	log.Fatal(err)

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
		return []byte(config.GetAuthfulConfig().Security.JwtKey), nil
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
