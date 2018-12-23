package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
	"io"
)

type Handler interface {
	Handle(c *gin.Context)
}

type CommonBody struct {
	Key string
	Val []byte
}

func readBody(b io.Reader) (*CommonBody, error) {
	res := new(CommonBody)

	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
