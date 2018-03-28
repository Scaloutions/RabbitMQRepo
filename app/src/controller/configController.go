package controller

import (
	"../service"
)

type (
	ConfigController struct {
		configService *service.ConfigService
	}
)

func GetConfigController() *ConfigController {
	return newConfigController()
}

func newConfigController() *ConfigController {
	configService := service.GetConfigService()
	return &ConfigController{configService}
}
