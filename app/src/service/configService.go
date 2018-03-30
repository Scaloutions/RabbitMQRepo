package service

import (
	"fmt"

	"github.com/streadway/amqp"

	"../config"
	"../consumer"
	"../producer"
	"../util"
)

type (
	ConfigService struct {
		u            *util.Utilities
		conn         *amqp.Connection
		consumerConn *amqp.Connection
	}
)

const (
	localHost = "localhost"
)

func GetConfigService() *ConfigService {
	return newConfigService()
}

func (configService ConfigService) GetUtilities() *util.Utilities {
	return configService.u
}

func (configService ConfigService) GetRabbitmqChannel() *amqp.Channel {

	conn := configService.conn
	if conn == nil { // when Rabbitmq is disabled
		return nil // TODO: add warning later
	}

	channel, err := conn.Channel()

	// TODO: add error handling later
	if err != nil {
		return nil
	}
	return channel
}

func (configService ConfigService) GetRabbitmqChannelForConsumer() *amqp.Channel {

	cConn := configService.consumerConn
	if cConn == nil { // when Rabbitmq is disabled
		return nil // TODO: add warning later
	}

	channel, err := cConn.Channel()

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

	queueConfig := configService.u.GetQueueConfig()

	queue, err1 := channel.QueueDeclare(
		queueConfig.Name,
		queueConfig.Durable,
		queueConfig.AutoDelete,
		queueConfig.Exclusive,
		queueConfig.NoWait,
		nil,
	)

	if err1 != nil {
		return nil
	}

	return &queue
}

func (configService ConfigService) GetSpecificQueue(index int) *amqp.Queue {

	channel := configService.GetRabbitmqChannel()
	queue := configService.getSpecificQueueHelper(channel, index)
	return queue
}

func (configService ConfigService) GetSpecificQueueForConsumer(
	index int) *amqp.Queue {

	channel := configService.GetRabbitmqChannelForConsumer()
	queue := configService.getSpecificQueueHelper(channel, index)
	return queue
}

func (configService ConfigService) getSpecificQueueHelper(channel *amqp.Channel, index int) *amqp.Queue {

	if channel == nil {
		return nil // TODO: add error handling later
	}

	queueConfig := configService.u.GetSpecificQueueConfig(index)

	queue, err1 := channel.QueueDeclare(
		queueConfig.Name,
		queueConfig.Durable,
		queueConfig.AutoDelete,
		queueConfig.Exclusive,
		queueConfig.NoWait,
		nil,
	)

	if err1 != nil {
		return nil
	}

	return &queue

}

func (configService ConfigService) CreateProducer(
	queue *amqp.Queue) *producer.Producer {

	return producer.NewProducer(queue)
}

func (configService ConfigService) CreateConsumer(
	queue *amqp.Queue) *consumer.Consumer {

	return consumer.NewConsumer(queue)
}

/*
	Private methods
*/

func newConfigService() *ConfigService {
	v := config.ReadInConfig()
	u := util.NewUtilites(v)
	conn := getRabbitmqConn(u, 0)
	cConn := getRabbitmqConn(u, 1)
	return &ConfigService{u, conn, cConn}
}

func getRabbitmqConn(u *util.Utilities, typeConn int) *amqp.Connection {

	if !u.IsRabbitmqConnEnabled() {
		//	TODO: add warning later
		fmt.Println("Rabbitmq is not enabled for current environment")
		return nil
	}

	var host string
	if typeConn == 0 { // connection for producer
		host = u.GetRabbitmqHost()
	} else { // connection for consumer
		host = localHost
	}

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
