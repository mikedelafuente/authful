package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weekendprojectapp/authful/users/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var userSvc *userService
var startTime time.Time

func init() {
	startTime = time.Now()
	fmt.Printf("Process started at %s\n", startTime)
	config.GetAuthfulConfig() // just attempt to get the config at startup
	config.GetDbConnection()  // just attempt to connect to the database at startup
}

func main() {
	userSvc = newUserService()

	myConfig := config.GetAuthfulConfig()
	fmt.Printf("\n\nAuthful: User Server running at %s:%v\n\n", myConfig.WebServer.Address, myConfig.WebServer.Port)
	setupRequestHandlers()
}

func setupRequestHandlers() {

	// Unsecured endpoints
	openR := myRouter.Methods(http.MethodGet, http.MethodPost).Subrouter()

	// ------------ UNPROTECTED API ENDPOINTS ------------
	// User signup/signin services
	openR.HandleFunc("/api/v1/users", userSvc.getUsers).Methods(http.MethodGet)
	openR.HandleFunc("/api/v1/users:signin", userSvc.authorizeUser).Methods(http.MethodPost)
	openR.HandleFunc("/api/v1/users:signup", userSvc.createUser).Methods(http.MethodPost)

	myConfig := config.GetAuthfulConfig()

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
