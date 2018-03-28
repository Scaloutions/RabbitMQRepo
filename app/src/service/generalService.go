package service

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"../data"
	"../util"
)

type (
	GeneralService struct {
		u *util.Utilities
	}
)

func GetGeneralService() *GeneralService {
	return newGeneralService()
}

func (generalService GeneralService) ProcessRequestBody(
	c *gin.Context) *data.Quote {

	var quote data.Quote

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil
	}

	util.UnserializeObject(body, &quote)

	return &quote
}

/*
	Private methods
*/

func newGeneralService() *GeneralService {

	configService := GetConfigService()
	u := configService.GetUtilities()
	return &GeneralService{u}
}
