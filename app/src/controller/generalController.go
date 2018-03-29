package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type (
	GeneralController struct {
		generalService *service.GeneralService
		configService  *service.ConfigService
		messageService *service.MessageService
		queue          *amqp.Queue
	}
)

func GetGeneralController() *GeneralController {
	return newGeneralController()
}

func (gController GeneralController) PublishMessage(c *gin.Context) {

	// get object from request
	body := gController.generalService.ProcessRequestBody(c)

	// publish message object in the queue
	producer := gController.configService.CreateProducer(
		gController.queue)
	channel := gController.configService.GetRabbitmqChannel()
	ch := gController.messageService.PublishMessage(
		body, channel, producer)
	<-ch
}

/*
	Private methods
*/

func newGeneralController() *GeneralController {

	gService := service.GetGeneralService()
	cService := service.GetConfigService()
	queue := cService.GetRabbitmqQueue()
	mService := service.GetMessageService(queue)

	return &GeneralController{gService, cService, mService, queue}
}
