package consumer

import "github.com/streadway/amqp"

type (
	Consumer struct {
		queue *amqp.Queue
	}
)

func NewConsumer(q *amqp.Queue) *Consumer {
	return &Consumer{q}
}
