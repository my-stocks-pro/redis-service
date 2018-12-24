package engine

import (
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/handler"
	"os"
)

type Service struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	redis    infrastructure.Redis
	handler  map[string]handler.Handler
	Engine   *gin.Engine
	QuitOS   chan os.Signal
	QuitRPC  chan bool
	QuitTick chan bool
}

func New(c infrastructure.Config, l infrastructure.Logger, r infrastructure.Redis) Service {
	return Service{
		config:   c,
		logger:   l,
		redis:    r,
		handler:  map[string]handler.Handler{},
		Engine:   gin.New(),
		QuitOS:   make(chan os.Signal),
		QuitRPC:  make(chan bool),
		QuitTick: make(chan bool),
	}
}
