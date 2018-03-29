package producer

import (
	"log"

	"github.com/streadway/amqp"
)

type (
	Producer struct {
		queue *amqp.Queue
	}
)

func NewProducer(q *amqp.Queue) *Producer {
	return &Producer{q}
}

func (producer Producer) Publish(
	channel *amqp.Channel, jsonObj []byte) {

	errC := make(chan error)

	go func() {
		err := channel.Publish(
			"",
			producer.queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonObj,
			})
		errC <- err
	}()

	err1 := <-errC

	if err1 != nil {
		// add error handling
		return
	}

	log.Println(string(jsonObj))
}
