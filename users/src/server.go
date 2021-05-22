package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"weekendproject.app/authful/users/config"
)

var myRouter = mux.NewRouter().StrictSlash(true)
var userSvc *userService

func init() {
	config.GetAuthfulConfig() // just attempt to get the config at startup
	config.GetDbConnection()  // just attempt to connect to the database at startup
}

func main() {
	userSvc = newUserService()

	myConfig := config.GetAuthfulConfig()
	fmt.Printf("Authful: User Server running at %s:%v", myConfig.WebServer.Address, myConfig.WebServer.Port)
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", myConfig.WebServer.Address, myConfig.WebServer.Port), myRouter))

}

func dbShutdown() {
	fmt.Println("shutting down database")
	db := config.GetDbConnection()
	db.Close()
}
