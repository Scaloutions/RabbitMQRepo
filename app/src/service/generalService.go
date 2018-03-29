package service

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type (
	GeneralService struct {
	}
)

func GetGeneralService() *GeneralService {
	return newGeneralService()
}

func (generalService GeneralService) ProcessRequestBody(
	c *gin.Context) []byte {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil
	}

	return body
}

/*
	Private methods
*/

func newGeneralService() *GeneralService {
	return &GeneralService{}
}
