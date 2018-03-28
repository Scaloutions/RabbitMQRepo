package service

import (
	"fmt"

	"github.com/streadway/amqp"

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

func (configService ConfigService) GetRabbitmqConn() *amqp.Connection {

	if !configService.u.IsRabbitmqConnEnabled() {
		//	TODO: add warning later
		fmt.Println("Rabbitmq is not enabled for current environment")
		return nil
	}

	host := configService.u.GetRabbitmqHost()
	port := configService.u.GetRabbitmqPort()
	connType := configService.u.GetRabbitmqConnType()
	pass := configService.u.GetRabbitmqPass()
	user := configService.u.GetRabbitmqUser()

	amqpURI := fmt.Sprintf("%s://%s:%s@%s:%d", connType, user, pass, host, port)

	// TODO: add error handling later
	conn := config.RabbitmqConnect(amqpURI)

	return conn
}
