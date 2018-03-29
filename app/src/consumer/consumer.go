package consumer

import (
	"log"

	"github.com/streadway/amqp"

	"./helper"
)

type (
	Consumer struct {
		queue *amqp.Queue
	}
)

func NewConsumer(q *amqp.Queue) *Consumer {
	return &Consumer{q}
}

func (consumer Consumer) Consume(channel *amqp.Channel) {

	msgs, err := channel.Consume(
		consumer.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			quote := helper.GetQuote(d.Body)
			log.Println("Quote: ", quote)

			// send quote to the cache server
		}
	}()

	// TODO: externalize message later
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
