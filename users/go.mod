module github.com/mikedelafuente/authful/users

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/mikedelafuente/authful-servertools v0.0.8
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

// replace github.com/mikedelafuente/authful-servertools => ../../authful-servertools
