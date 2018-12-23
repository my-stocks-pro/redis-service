package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
	"fmt"
	"errors"
)

type Default interface {
	Handle(c *gin.Context)
}

type DefaultType struct {
	config infrastructure.Config
	logger infrastructure.Logger
	redis  infrastructure.Redis
}

func NewDefault(l infrastructure.Logger) DefaultType {
	return DefaultType{
		logger: l,
	}
}

func (d DefaultType) Handle(c *gin.Context) {
	paramKey := "service"
	keyRedisDB := c.Param(paramKey)
	if keyRedisDB == "" {
		d.logger.ContextError(c, http.StatusNotAcceptable, errors.New(fmt.Sprintf("Key -> %s dont exist in gin Params", keyRedisDB)))
		return
	}
	d.logger.ContextError(c, http.StatusNotAcceptable, errors.New(fmt.Sprintf("Key: %s Not Acceptable", keyRedisDB)))
}
