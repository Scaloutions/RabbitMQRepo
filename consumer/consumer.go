package main

import (
	"log"

	"github.com/streadway/amqp"

	"../app/src/config"
	"../app/src/consumer"
	"../app/src/util"

	"./worker"
)

const (
	amqpURI = "amqp://guest:guest@localhost:5672/"
)

func createConnection() *amqp.Connection {

	c := make(chan *amqp.Connection)

	go func() {
		conn, err := amqp.Dial(amqpURI)
		if err != nil {
			c <- nil
		} else {
			c <- conn
		}
	}()

	conn := <-c
	return conn
}

func createChannel(conn *amqp.Connection) *amqp.Channel {

	c := make(chan *amqp.Channel)

	go func() {
		ch, err := conn.Channel()
		if err != nil {
			c <- nil
		} else {
			c <- ch
		}
	}()

	ch := <-c
	return ch
}

func createQueue(ch *amqp.Channel) *amqp.Queue {

	v := config.ReadInConfig()
	u := util.NewUtilites(v)
	queueConfig := u.GetQueueConfig()

	queue, err := ch.QueueDeclare(
		queueConfig.Name,
		queueConfig.Durable,
		queueConfig.AutoDelete,
		queueConfig.Exclusive,
		queueConfig.NoWait,
		nil,
	)

	if err != nil {
		return nil
	}

	return &queue
}

func mainHelper() {
	conn := createConnection()
	if conn == nil {
		log.Fatalln(conn)
	}

	c := make(chan int)

	go func() {
		ch := createChannel(conn)
		if ch == nil {
			log.Fatalln(ch)
		}
		defer ch.Close()
		queue := createQueue(ch)
		consumer := consumer.NewConsumer(queue)
		consumer.Consume(ch)
		c <- 0
	}()

	<-c

}

func main() {

	// saveWorker := worker.GetSaveWorker()
	// saveWorker.ConsumeMessage()

	getByKeyWorker := worker.GetGetByKeyWorker()
	getByKeyWorker.ConsumeMessage()

	// getWorker := worker.GetGetWorker()
	// go getWorker.ConsumeMessage()

}
