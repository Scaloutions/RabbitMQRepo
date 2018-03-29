package helper

import (
	"../../data"
	"../../util"
)

func GetQuote(body []byte) *data.Quote {

	var quote *data.Quote
	util.UnserializeObject(body, &quote)
	return quote
}
