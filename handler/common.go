package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
	"github.com/pkg/errors"
)

type Common interface {
	Get(c *gin.Context) error
	Post(c *gin.Context) error
	Put(c *gin.Context) error
	Delete(c *gin.Context) error
}

type TypeCommon struct {
	config infrastructure.Config
	logger infrastructure.Logger
	redis  infrastructure.Redis
}

func NewCommon(c infrastructure.Config, l infrastructure.Logger, r infrastructure.Redis) TypeCommon {
	return TypeCommon{
		config: c,
		logger: l,
		redis:  r,
	}
}

func (s TypeCommon) Get(c *gin.Context) error {
	serviceName := "service"
	serviceType := c.Param(serviceName)

	db, err := s.redis.GetDB(serviceType, serviceName)
	if err != nil {
		s.logger.ContextError(c, http.StatusInternalServerError, err)
		return
	}

	body, err := readBody(c.Request.Body)
	if err != nil {
		s.logger.ContextError(c, http.StatusInternalServerError, err)
		return
	}


	switch c.Request.Method {
	case http.MethodGet:
		val, err := s.redis.Get(body.Key, db)
		if err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

		_, err = c.Writer.Write(val)
		if err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	case http.MethodPost:
		if err := s.redis.Set(body.Key, body.Val, db); err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	case http.MethodDelete:
		err := s.redis.Delete(body.Key, db)
		if err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	default:
		s.logger.ContextError(c, http.StatusMethodNotAllowed, errors.New("Method Not Allowed"))
		return
	}

	s.logger.ContextSuccess(c, http.StatusOK)
}


func (v TypeCommon) Post(c *gin.Context) error {

	return nil
}

func (v TypeCommon) Put(c *gin.Context) error {

	return nil
}

func (v TypeCommon) Delete(c *gin.Context) error {

	return nil
}