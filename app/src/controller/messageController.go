package controller

import (
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

func GetMessageController() *MessageController {
	return newMessageController()
}

func (mController MessageController) ConsumeMessage() {

	consumer := mController.configService.CreateConsumer(
		mController.queue)
	channel := mController.configService.GetRabbitmqChannel()
	ch := mController.messageService.ConsumeMessage(
		channel, consumer)
	<-ch
}

/*
	Private methods
*/

func newMessageController() *MessageController {

	gService := service.GetGeneralService()
	cService := service.GetConfigService()
	queue := cService.GetRabbitmqQueue()
	mService := service.GetMessageService(queue)

	return &MessageController{gService, cService, mService, queue}
}
