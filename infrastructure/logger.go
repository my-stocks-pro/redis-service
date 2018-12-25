package infrastructure

import (
	"fmt"
	"os"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
)

const (
	logPath   = "app_log"
	logPrefix = "redis-service"
)

type Logger interface {
	Error(msg string)
	Info(msg string)
	ContextError(c *gin.Context, status int, err error)
	ContextSuccess(c *gin.Context, status int)
}

type LoggerType struct {
	Client *zap.Logger
}

func NewLogger() (LoggerType, error) {
	filename := fmt.Sprintf("%s/%s.log", logPath, logPrefix)
	_, err := os.Create(filename)
	if err != nil {
		return LoggerType{}, err
	}

	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}

	zapLog, err := conf.Build()
	if err != nil {
		return LoggerType{}, err
	}

	defer zapLog.Sync()

	return LoggerType{Client: zapLog}, nil
}

func (l LoggerType) Error(msg string) {
	fmt.Println(msg)
	//l.Client.Error(msg)
}

func (l LoggerType) Info(msg string) {
	fmt.Println(msg)
	//l.Client.Info(msg)
}

func (l LoggerType) ContextError(c *gin.Context, status int, err error) {
	l.Error(fmt.Sprintf("STATUS: %d, Message: %s", status, err.Error()))
	c.JSON(status, gin.H{
		"status": status,
		"error":  err.Error(),
	})
}

func (l LoggerType) ContextSuccess(c *gin.Context, status int) {
	l.Info(fmt.Sprintf("STATUS: %d, ErrorMessage: NIL", status))
	c.JSON(status, gin.H{
		"status": status,
		"error":  nil,
	})
}
