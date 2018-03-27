FROM golang:1.9.2-alpine3.6 AS build

# Install tools required to build the project
RUN apk add --no-cache git

RUN go get github.com/streadway/amqp