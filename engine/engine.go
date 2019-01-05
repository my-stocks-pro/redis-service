package engine

import (
	"github.com/my-stocks-pro/redis-service/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

const (
	version   = "version"
	health    = "health"
	earnings  = "earnings"
	approved  = "approved"
	rejected  = "rejected"
	pinterest = "pinterest"
)

func (s *Service) InitMux() {
	s.Engine.GET("/:service", s.HandleFunc)
	s.Engine.POST("/:service", s.HandleFunc)
	s.Engine.PUT("/:service", s.HandleFunc)
	s.Engine.DELETE("/:service", s.HandleFunc)
}

func (s *Service) getHandler(serviceName string) handler.Handler {
	_, ok := s.handler[serviceName]
	if !ok {
		s.handler[serviceName] = s.HandlerConstruct(serviceName)
	}
	return s.handler[serviceName]
}

func (s *Service) HandlerConstruct(serviceName string) handler.Handler {
	switch serviceName {
	case version:
		return handler.NewVersion(s.config)
	case health:
		return handler.NewHealth(s.config, s.redis)
	case earnings, approved, rejected:
		return handler.NewCommon(s.config, s.logger, s.redis)
	case pinterest:
		return handler.NewPinterest(s.config, s.logger, s.redis)
	default:
		return nil
	}
	return nil
}

func (s *Service) HandleFunc(c *gin.Context) {
	serviceType := c.Param("service")

	h := s.getHandler(serviceType)
	if h == nil {
		//TODO
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		if err := h.Get(c); err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := h.Post(c); err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPut:
		if err := h.Put(c); err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	case http.MethodDelete:
		if err := h.Delete(c); err != nil {
			s.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	default:
		s.logger.ContextError(c, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	s.logger.ContextSuccess(c, http.StatusOK)

}
