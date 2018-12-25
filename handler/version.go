package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
	"time"
)

type Version interface {
	Handle(c *gin.Context)
}

type TypeVersion struct {
	config infrastructure.Config
}

func NewVersion(c infrastructure.Config) TypeVersion {
	return TypeVersion{
		config: c,
	}
}

func (v TypeVersion) Handle(c *gin.Context) {

	switch c.Request.Method {
	case http.MethodGet:
		c.JSON(http.StatusOK, gin.H{
			"startTime": v.config.StartDate,
			"currDate":  time.Now().Format("2006-01-02 15:04"),
			"version":   "1.0",
			"service":   v.config.Name,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
	}

}
