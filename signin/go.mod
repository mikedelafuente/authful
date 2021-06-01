module github.com/mikedelafuente/authful/signin

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/mikedelafuente/authful-servertools v0.0.8
)

// replace github.com/mikedelafuente/authful-servertools => ../../authful-servertools
