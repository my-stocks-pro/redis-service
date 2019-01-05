package engine

import (
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

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

func (p Service) getParams(c *gin.Context) (int, *CommonBody, error) {
	serviceName := "service"
	serviceType := c.Param(serviceName)

	db, err := p.redis.GetDB(serviceType, serviceName)
	if err != nil {
		return 0, nil, err
	}

	body, err := readBody(c.Request.Body)
	if err != nil {
		return 0, nil, err
	}
	return db, body, nil
}
