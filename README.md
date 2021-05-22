# oauth-server-golang
A sample OAuth server written in Go

# TODO
- Use conversion for objects instead mapping - determine the Golang way of doing this
- Look into using configuration to set keys and JWT timeouts
- Create a common approach for returning errors

# Endpoints

## Homepage
- ANY /

## User signup/signin services
- POST /auth/signin
- POST /auth/signup

## Developer services
- POST /dev/signin
- POST /dev/signup

## Authorization Code Flow
- GET /oauth/authorize
- POST /oauth/authorize
- POST /oauth/token

## User services
- GET /api/v1/users
- GET /api/v1/users/{id}
- GET /api/v1/users:byUsername/{username}

## Developer services
- POST /api/v1/dev/my/keys
- GET /api/v1/dev/my/keys
- DELETE /api/v1/dev/my/keys/{client_id}