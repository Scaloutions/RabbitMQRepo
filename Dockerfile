FROM golang:1.9.2-alpine3.6 AS build

# Install tools required to build the project
RUN apk add --no-cache git

RUN go get github.com/streadway/amqp

RUN go get "github.com/gin-gonic/gin"

RUN go get -v "github.com/spf13/viper"

RUN mkdir -p /app

ADD . /app 

WORKDIR /app

RUN go build ./server.go

CMD ["./server"]