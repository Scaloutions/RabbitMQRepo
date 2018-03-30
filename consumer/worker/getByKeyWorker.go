package worker

import (
	"../../app/src/controller"
)

type GetByKeyWorker struct {
	mController *controller.MessageController
}

func GetGetByKeyWorker() *GetByKeyWorker {
	return newGetByKeyWorker()
}

func (gBKWorker GetByKeyWorker) ConsumeMessage() {

	c := make(chan int)
	go func() {
		consumer, channel := gBKWorker.mController.GetConsumerAndChannel()
		consumer.Consume(channel)
		c <- 0
	}()

	<-c
}

/*
	Private methods
*/

func newGetByKeyWorker() *GetByKeyWorker {
	mController := controller.GetMessageController(1)
	return &GetByKeyWorker{mController}
}
