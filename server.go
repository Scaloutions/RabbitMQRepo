package main

import (
	"./app/src/service"
)

func main() {

	configService := service.GetConfigService()

	// config.RabbitmqConnect("amqp://guest:guest@rabbitmq:5672")
	configService.GetRabbitmqConn()
}
