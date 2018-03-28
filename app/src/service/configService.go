package service

import (
	"../config"
	"../util"
)

type (
	ConfigService struct {
		u *util.Utilities
	}
)

func GetConfigService() *ConfigService {
	return newConfigService()
}

func newConfigService() *ConfigService {
	v := config.ReadInConfig()
	u := util.NewUtilites(v)
	return &ConfigService{u}
}

func (configService ConfigService) GetUtilities() *util.Utilities {
	return configService.u
}
