package main

import (
	"fmt"
	"log"

	"./app/src/controller"
	"github.com/gin-gonic/gin"
)

func testing(c *gin.Context) {
	log.Println("Hello World")
}

func getMainEngine() (*gin.Engine, string) {
	router := gin.Default()
	configController := controller.GetConfigController()
	generalController := controller.GetGeneralController()
	u := configController.GetUtilites()
	api := router.Group(u.GetStringConfigValue("api.api"))
	{
		api.POST("/testing", testing)
		api.POST(
			u.GetStringConfigValue("api.publish"),
			generalController.PublishMessage)
	}

	portStr := fmt.Sprintf(":%d", u.GetIntConfigValue("general.port"))

	return router, portStr
}

func setUp() {
	router, portStr := getMainEngine()
	router.Run(portStr)
}

func main() {
	setUp()
}
