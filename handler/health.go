package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
)

type Health interface {
	Get(c *gin.Context) error
	Post(c *gin.Context) error
	Put(c *gin.Context) error
	Delete(c *gin.Context) error
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

func (h HealthType) Get(c *gin.Context) error {

	switch c.Request.Method {
	case http.MethodGet:
		redisStatus := true
		if err := h.redis.Ping(); err != nil {
			redisStatus = false
		}

		c.JSON(http.StatusOK, gin.H{
			"service": true,
			"redisDB": redisStatus,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
	}

}

func (h HealthType) Post(c *gin.Context) error {

	return nil
}

func (h HealthType) Put(c *gin.Context) error {

	return nil
}

func (h HealthType) Delete(c *gin.Context) error {

	return nil
}
