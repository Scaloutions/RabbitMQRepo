package main

import (
	"fmt"

	"./app/src/service"
)

func main() {

	configService := service.GetConfigService()
	u := configService.GetUtilities()
	fmt.Println(u.GetStringConfigValue("general.rabbitmq.connection_type"))
}
