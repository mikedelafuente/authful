package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/users/internal/config"
	"github.com/mikedelafuente/authful/users/internal/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	log.SetOutput(os.Stdout)
	startTime = time.Now()
	config.GetConfig()
	config.GetDbConnection() // just attempt to connect to the database at startup
	fmt.Printf("Process started at %s\n", startTime)
}

func main() {
	logger.Printf("\n\nAuthful: User Server running at %s:%v\n\n", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port)
	logger.Printf("Log Level: %s\n", logger.GetLogLevel())
	setupRequestHandlers()
}

func setupRequestHandlers() {

	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost).Subrouter()
	openR.HandleFunc("/api/v1/account:signin", controllers.AccountSigninPost).Methods(http.MethodPost)
	openR.HandleFunc("/api/v1/account:signup", controllers.AccountSignupPost).Methods(http.MethodPost)
	openR.Use(openHandler)

	// ------------ UNPROTECTED API ENDPOINTS ------------
	// User signup/signin services
	secureUserR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	secureUserR.HandleFunc("/api/v1/users", controllers.UsersGet).Methods(http.MethodGet)
	secureUserR.Use(bearerJwtHandler)

	defer dbShutdown()
	err := http.ListenAndServe(fmt.Sprintf("%s:%v", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port), myRouter)
	endTime := time.Now()
	logger.Verbose(context.Background(), fmt.Sprintf("Process stopped at %s\n", endTime))
	elapsed := endTime.Sub(startTime)
	logger.Verbose(context.Background(), fmt.Sprintf("Server uptime was: %s", elapsed))
	logger.Fatal(context.Background(), err)
}

func dbShutdown() {
	logger.Verbose(context.Background(), "shutting down database")
	db := config.GetDbConnection()
	db.Close()
}

func openHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = extractAndSetTraceId(r)
		cookie, err := r.Cookie("userSessionToken")
		if err == nil {
			// Redirect
			rawToken := cookie.Value
			_, r = processToken(rawToken, r)
		}
		logger.Verbose(r.Context(), fmt.Sprintf("Request recieved: %s %s", r.Method, r.URL))
		next.ServeHTTP(w, r)
	})
}

func bearerJwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Values("Authorization")
		r = extractAndSetTraceId(r)
		isValid := false

		if len(authHeader) > 0 {
			if strings.HasPrefix(authHeader[0], "Bearer ") {
				parts := strings.Split(authHeader[0], " ")
				rawToken := parts[1]

				isValid, r = processToken(rawToken, r)
			}
		}

		logger.Verbose(r.Context(), fmt.Sprintf("Request recieved: %s %s", r.Method, r.URL))
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func extractAndSetTraceId(r *http.Request) *http.Request {
	traceIdParts := r.Header.Values("x-trace-id")

	var traceId string
	if len(traceIdParts) == 0 || len(traceIdParts[0]) == 0 {
		traceId = uuid.New().String()
	} else {
		traceId = traceIdParts[0]
	}
	ctx := context.WithValue(r.Context(), customclaims.ContextTraceId, traceId)
	r = r.WithContext(ctx)
	return r
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
		logger.Error(r.Context(), err)
	}

	ctx := context.WithValue(r.Context(), customclaims.ContextKeyUserId, userId)
	ctx = context.WithValue(ctx, customclaims.ContextJwt, token.Raw)
	r = r.WithContext(ctx)

	return isValid, r
}
