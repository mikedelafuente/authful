module github.com/weekendprojectapp/authful/signin

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/weekendprojectapp/authful/server v0.0.0
)

replace github.com/weekendprojectapp/authful/server => ../server
