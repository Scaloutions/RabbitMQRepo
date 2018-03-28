package main

import (
	"fmt"
	"log"

	"./app/src/controller"
	"./app/src/service"
	"github.com/gin-gonic/gin"
)

func testing(c *gin.Context) {
	generalService := service.GetGeneralService()
	quote := generalService.ProcessRequestBody(c)
	log.Println(*quote)
}

func getMainEngine() (*gin.Engine, string) {
	router := gin.Default()
	configController := controller.GetConfigController()
	u := configController.GetUtilites()
	api := router.Group(u.GetStringConfigValue("api.api"))
	{
		// api.GET("/testing", testing)
		api.POST("/testing", testing)
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
