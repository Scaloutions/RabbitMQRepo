package controller

import (
	"../service"
	"../util"
)

type (
	ConfigController struct {
		configService *service.ConfigService
	}
)

func GetConfigController() *ConfigController {
	return newConfigController()
}

func (configController ConfigController) GetUtilites() *util.Utilities {
	return configController.configService.GetUtilities()
}

/*
	Private methods
*/

func newConfigController() *ConfigController {
	configService := service.GetConfigService()
	return &ConfigController{configService}
}
