package worker

import (
	"../../app/src/controller"
)

type GetWorker struct {
	mController *controller.MessageController
}

func GetGetWorker() *GetWorker {
	return newGetWorker()
}

func (gWorker GetWorker) ConsumeMessage() {

	c := make(chan int)
	go func() {
		consumer, channel := gWorker.mController.GetConsumerAndChannel()
		consumer.Consume(channel)
		c <- 0
	}()

	<-c
}

/*
	Private methods
*/

func newGetWorker() *GetWorker {
	mController := controller.GetMessageController(2)
	return &GetWorker{mController}
}
