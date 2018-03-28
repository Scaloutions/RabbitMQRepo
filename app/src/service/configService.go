package service

import (
	"fmt"

	"github.com/streadway/amqp"

	"../config"
	"../util"
)

type (
	ConfigService struct {
		u    *util.Utilities
		conn *amqp.Connection
	}
)

func GetConfigService() *ConfigService {
	return newConfigService()
}

func (configService ConfigService) GetRabbitmqChannel() *amqp.Channel {

	conn := configService.conn
	if conn == nil { // when Rabbitmq is disabled
		return nil // TODO: add warning later
	}

	channel, err := configService.conn.Channel()

	// TODO: add error handling later
	if err != nil {
		return nil
	}
	return channel
}

func (configService ConfigService) GetRabbitmqQueue() *amqp.Queue {

	channel := configService.GetRabbitmqChannel()
	if channel == nil {
		return nil // TODO: add error handling later
	}

	queue, err1 := channel.QueueDeclare(

		"retaildash-scrapy", // name

		false, // durable

		false, // delete when unused

		false, // exclusive

		false, // no-wait (wait time for processing)

		nil, // arguments

	)
	if err1 != nil {
		return nil
	}

	return &queue
}

/*
	Private methods
*/

func newConfigService() *ConfigService {
	v := config.ReadInConfig()
	u := util.NewUtilites(v)
	conn := getRabbitmqConn(u)
	return &ConfigService{u, conn}
}

func getRabbitmqConn(u *util.Utilities) *amqp.Connection {

	if !u.IsRabbitmqConnEnabled() {
		//	TODO: add warning later
		fmt.Println("Rabbitmq is not enabled for current environment")
		return nil
	}

	host := u.GetRabbitmqHost()
	port := u.GetRabbitmqPort()
	connType := u.GetRabbitmqConnType()
	pass := u.GetRabbitmqPass()
	user := u.GetRabbitmqUser()

	amqpURI := fmt.Sprintf("%s://%s:%s@%s:%d", connType, user, pass, host, port)

	// TODO: add error handling later

	c := make(chan *amqp.Connection)

	go func() {
		conn := config.RabbitmqConnect(amqpURI)
		c <- conn
	}()

	conn := <-c

	return conn
}
