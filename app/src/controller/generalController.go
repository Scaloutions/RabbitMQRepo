package controller

import (
	"../producer"
	"../service"
	"../util"
	"github.com/gin-gonic/gin"
)

type (
	GeneralController struct {
		generalService *service.GeneralService
		configService  *service.ConfigService
		messageService *service.MessageService
	}
)

func GetGeneralController() *GeneralController {
	return newGeneralController()
}

func (gController GeneralController) PublishMessage(c *gin.Context) {

	// get object from request
	reqBody := gController.generalService.ProcessRequestBody(c)
	requestType := reqBody.RequestType
	body := util.SerializeObject(reqBody.QuoteObj)

	// // publish message object in the queue
	producer := gController.createProducerWithSpecifiedQueue(requestType)
	channel := gController.configService.GetRabbitmqChannel()
	ch := gController.messageService.PublishMessage(
		body, channel, producer)
	<-ch
}

/*
	Private methods
*/

func (gController GeneralController) createProducerWithSpecifiedQueue(
	index int) *producer.Producer {

	queue := gController.configService.GetSpecificQueue(index)
	producer := gController.configService.CreateProducer(queue)

	return producer
}

func newGeneralController() *GeneralController {

	gService := service.GetGeneralService()
	cService := service.GetConfigService()
	mService := service.GetMessageService()

	return &GeneralController{gService, cService, mService}
}
