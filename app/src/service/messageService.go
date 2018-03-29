package service

import (
	"github.com/streadway/amqp"

	"../consumer"
	"../producer"
)

type (
	MessageService struct {
		queue *amqp.Queue
	}
)

func GetMessageService(q *amqp.Queue) *MessageService {
	return newMessageService(q)
}

func (messageService MessageService) PublishMessage(
	jsonObject []byte,
	channel *amqp.Channel,
	producer *producer.Producer) <-chan int {

	c := make(chan int)

	go func() {
		producer.Publish(channel, jsonObject)
		c <- 0
	}()

	return c
}

func (messageService MessageService) ConsumeMessage(
	channel *amqp.Channel,
	consumer *consumer.Consumer) <-chan int {

	c := make(chan int)

	go func() {
		consumer.Consume(channel)
		c <- 0
	}()

	return c
}

/*
	Private methods
*/

func newMessageService(q *amqp.Queue) *MessageService {

	return &MessageService{q}
}
