package engine

import (
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/handler"
)

type Server struct {
	config  infrastructure.Config
	logger  infrastructure.Logger
	redis   infrastructure.Redis
	handler map[string]handler.Handler
	Engine  *gin.Engine
}

func New(c infrastructure.Config, l infrastructure.Logger, r infrastructure.Redis) Server {
	return Server{
		config:  c,
		logger:  l,
		redis:   r,
		handler: map[string]handler.Handler{},
		Engine:  gin.New(),
	}
}
