#copy stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src
COPY . .

#build stage
# WORKDIR /go/src/users
# RUN go get -d -v ./...
# RUN go build -o /go/bin/users -v ./main.go
# RUN cp -r /go/src/users/config.json /go/bin/users

#final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/users  /users
# ENTRYPOINT /users
# LABEL Name=users Version=0.0.1
# EXPOSE 8081

WORKDIR /go/src/users
RUN go install .
RUN cp -r /go/src/users/config.json /go/bin/users
ENTRYPOINT /go/bin/users

EXPOSE 8081