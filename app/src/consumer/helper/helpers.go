package helper

import (
	"../../config"
	"../../data"
	"../../util"
)

type Helper struct {
	u *util.Utilities
}

func GetHelper() *Helper {
	return newHelper()
}

func GetQuote(body []byte) *data.Quote {

	var quote *data.Quote
	util.UnserializeObject(body, &quote)
	return quote
}

func ProcessRequestBody(body []byte) *data.RequestBody {

	var reqBody *data.RequestBody
	util.UnserializeObject(body, &reqBody)
	return reqBody
}

func (helper Helper) SendRequest(body []byte, index int) {

	c := helper.u.SendPostRequest(body, index)
	<-c
}

/*
	Private methods
*/

func newHelper() *Helper {
	v := config.ReadInConfig()
	u := util.NewUtilites(v)
	return &Helper{u}
}
