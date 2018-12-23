package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
)

type Health interface {
	Handle(c *gin.Context)
}

type HealthType struct {
	config infrastructure.Config
	redis  infrastructure.Redis
}

func NewHealth(c infrastructure.Config, r infrastructure.Redis) HealthType {
	return HealthType{
		config: c,
		redis:  r,
	}
}

func (h HealthType) Handle(c *gin.Context) {
	redisStatus := true
	if err := h.redis.Ping(); err != nil {
		redisStatus = false
	}

	c.JSON(200, gin.H{
		"service": true,
		"redisDB": redisStatus,
	})
}
