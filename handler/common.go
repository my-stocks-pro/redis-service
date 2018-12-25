package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
)

type Common interface {
	Handle(c *gin.Context)
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

func (s TypeCommon) Handle(c *gin.Context) {
	serviceName := "service"
	db, err := s.redis.GetDB(c.Param(serviceName), serviceName)
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
		if err = s.redis.Set(body.Key, body.Val, db); err != nil {
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
		c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	s.logger.ContextSuccess(c, http.StatusOK, nil)
}
