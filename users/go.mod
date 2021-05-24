module github.com/weekendprojectapp/authful/users

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/weekendprojectapp/authful/serverutils v0.0.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

replace github.com/weekendprojectapp/authful/serverutils => ../serverutils

