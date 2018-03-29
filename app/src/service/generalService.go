package service

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"../data"
	"../util"
)

type (
	GeneralService struct {
	}
)

func GetGeneralService() *GeneralService {
	return newGeneralService()
}

func (generalService GeneralService) ProcessRequestBody(
	c *gin.Context) *data.RequestBody {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil
	}

	var reqBody *data.RequestBody
	util.UnserializeObject(body, &reqBody)

	return reqBody
}

/*
	Private methods
*/

func newGeneralService() *GeneralService {
	return &GeneralService{}
}
