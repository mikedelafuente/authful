#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/settings/config.json /settings/config.json
COPY --from=builder /go/src/app/settings/proxy.json /settings/proxy.json
ENTRYPOINT /app
LABEL Name=authful-proxy Version=0.0.1
EXPOSE 8080
