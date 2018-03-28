package main

import (
	"fmt"

	"./app/src/service"
)

func main() {

	configService := service.GetConfigService()

	channel := configService.GetRabbitmqChannel()
	fmt.Println(channel)

	queue := configService.GetRabbitmqQueue()
	fmt.Println(queue)

}
