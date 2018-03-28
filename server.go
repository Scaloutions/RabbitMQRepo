package main

import (
	"./app/src/config"
)

func main() {

	// configService := service.GetConfigService()
	// u := configService.GetUtilities()
	// fmt.Println(u.GetStringConfigValue("general.rabbitmq.connection_type"))

	config.RabbitmqConnect()
}
