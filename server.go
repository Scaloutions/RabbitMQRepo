package main

import (
	"fmt"

	"./app/src/controller"
	"github.com/gin-gonic/gin"
)

func testing(c *gin.Context) {
	fmt.Println("Hello World")
}

func getMainEngine() (*gin.Engine, string) {
	router := gin.Default()
	configController := controller.GetConfigController()
	u := configController.GetUtilites()
	api := router.Group(u.GetStringConfigValue("api.api"))
	{
		api.GET("/testing", testing)
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
