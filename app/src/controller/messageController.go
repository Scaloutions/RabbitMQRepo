package controller

import (
	"../consumer"
	"../service"
	"github.com/streadway/amqp"
)

type (
	MessageController struct {
		generalService *service.GeneralService
		configService  *service.ConfigService
		messageService *service.MessageService
		queue          *amqp.Queue
	}
)

func GetMessageController(index int) *MessageController {
	return newMessageController(index)
}

func (mController MessageController) ConsumeMessage() {

	consumer := mController.createConsumerWithSpecifiedQueue()
	channel := mController.configService.GetRabbitmqChannelForConsumer()
	ch := mController.messageService.ConsumeMessage(channel, consumer)
	<-ch
}

func (mController MessageController) GetConsumerAndChannel() (
	*consumer.Consumer, *amqp.Channel) {
	consumer := mController.createConsumerWithSpecifiedQueue()
	channel := mController.configService.GetRabbitmqChannelForConsumer()
	return consumer, channel
}

/*
	Private methods
*/

func (mController MessageController) createConsumerWithSpecifiedQueue() *consumer.Consumer {

	queue := mController.queue
	consumer := mController.configService.CreateConsumer(queue)

	return consumer
}

func newMessageController(index int) *MessageController {

	gService := service.GetGeneralService()
	cService := service.GetConfigService()
	mService := service.GetMessageService()
	queue := cService.GetSpecificQueueForConsumer(index)

	return &MessageController{gService, cService, mService, queue}
}
