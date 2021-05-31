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
	"github.com/google/uuid"
	"github.com/mikedelafuente/authful-servertools/pkg/customclaims"
	"github.com/mikedelafuente/authful-servertools/pkg/logger"
	"github.com/mikedelafuente/authful/signin/internal/config"
	"github.com/mikedelafuente/authful/signin/internal/controllers"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var startTime time.Time

func init() {
	log.SetOutput(os.Stdout)
	startTime = time.Now()
	config.GetConfig()
	fmt.Printf("Process started at %s\n", startTime)
}

func main() {
	fmt.Printf("\n\nAuthful: Signin Server running at %s:%v\n\n", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port)
	fmt.Printf("Log Level: %s\n", logger.GetLogLevel())
	setupRequestHandlers()
}

func setupRequestHandlers() {
	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodPut).Subrouter()
	openR.HandleFunc("/login", controllers.DisplayLogin).Methods(http.MethodGet)
	openR.HandleFunc("/login", controllers.ProcessLogin).Methods(http.MethodPost)
	openR.HandleFunc("/signup", controllers.DisplaySignup).Methods(http.MethodGet)
	openR.HandleFunc("/signup", controllers.ProcessSignup).Methods(http.MethodPost)
	openR.Use(openHandler)

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

	err := http.ListenAndServe(fmt.Sprintf("%s:%v", config.GetConfig().WebServer.Host, config.GetConfig().WebServer.Port), myRouter)

	endTime := time.Now()
	logger.Verbose(context.Background(), fmt.Sprintf("Process stopped at %s\n", endTime))
	elapsed := endTime.Sub(startTime)
	logger.Verbose(context.Background(), fmt.Sprintf("Server uptime was: %s", elapsed))
	logger.Fatal(context.Background(), err)

}

func openHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = extractAndSetTraceId(r)

		next.ServeHTTP(w, r)
	})
}

func cookieJwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isValid := true
		r = extractAndSetTraceId(r)

		cookie, err := r.Cookie("userSessionToken")
		if err != nil {
			logger.Error(r.Context(), err)
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
			logger.Debug(r.Context(), "Valid token passed")
		}
	} else {
		logger.Error(r.Context(), err)
	}

	ctx := context.WithValue(r.Context(), customclaims.ContextKeyUserId, userId)
	ctx = context.WithValue(ctx, customclaims.ContextJwt, token.Raw)
	r = r.WithContext(ctx)

	return isValid, r
}
