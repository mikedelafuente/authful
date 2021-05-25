module github.com/weekendprojectapp/authful/signin

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/weekendprojectapp/authful/serverutils v0.0.0
	github.com/weekendprojectapp/authful/users v0.0.0-20210524153354-a82d9d39db8c
)

replace github.com/weekendprojectapp/authful/serverutils => ../serverutils
