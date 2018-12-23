package engine

import (
	"github.com/my-stocks-pro/redis-service/handler"
	"github.com/gin-gonic/gin"
)

const (
	version  = "version"
	health   = "health"
	earnings = "earnings"
	approved = "approved"
	rejected = "rejected"
)

func (s *Server) InitMux() {
	s.Engine.GET("/:service", s.GetHandler)
	s.Engine.POST("/:service", s.GetHandler)
}

func (s *Server) GetHandler(c *gin.Context) {
	serviceType := c.Param("service")
	_, ok := s.handler[serviceType]
	if !ok {
		s.handler[serviceType] = s.HandlerConstruct(serviceType)
	}
	s.handler[serviceType].Handle(c)
}

func (s *Server) HandlerConstruct(serviceType string) handler.Handler {
	switch serviceType {
	case version:
		return handler.NewVersion(s.config)
	case health:
		return handler.NewHealth(s.config, s.redis)
	case earnings, approved, rejected:
		return handler.NewCommon(s.config, s.logger, s.redis)
	default:
		return handler.NewDefault(s.logger)
	}
	return nil
}


