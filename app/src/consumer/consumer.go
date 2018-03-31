package consumer

import (
	"log"

	"../util"
	"./helper"
	"github.com/streadway/amqp"
)

type (
	Consumer struct {
		queue  *amqp.Queue
		helper *helper.Helper
	}
)

func NewConsumer(q *amqp.Queue) *Consumer {
	helper := helper.GetHelper()
	return &Consumer{q, helper}
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
			reqBody := helper.ProcessRequestBody(d.Body)
			log.Println("Request: ", reqBody)
			requestType := reqBody.RequestType
			body := util.SerializeObject(reqBody.QuoteObj)

			// send quote to the cache server
			consumer.helper.SendRequest(body, requestType)
		}

	}()

	// TODO: externalize message later
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
