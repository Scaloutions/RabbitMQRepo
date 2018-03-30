package worker

import (
	"../../app/src/controller"
)

type SaveWorker struct {
	mController *controller.MessageController
}

func GetSaveWorker() *SaveWorker {
	return newSaveWorker()
}

func (saveWorker SaveWorker) ConsumeMessage() {

	c := make(chan int)
	go func() {
		consumer, channel := saveWorker.mController.GetConsumerAndChannel()
		consumer.Consume(channel)
		c <- 0
	}()

	<-c
}

/*
	Private methods
*/

func newSaveWorker() *SaveWorker {
	mController := controller.GetMessageController(0)
	return &SaveWorker{mController}
}
