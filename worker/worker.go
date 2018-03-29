package main

import (
	"../app/src/controller"
)

func main() {

	messageController := controller.GetMessageController()
	ch := make(chan int)

	go func() {
		messageController.ConsumeMessage()
		ch <- 0
	}()

	<-ch

}
