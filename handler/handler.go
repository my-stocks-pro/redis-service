package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	//Handle(c *gin.Context)
	Get(c *gin.Context) error
	Post(c *gin.Context) error
	Put(c *gin.Context) error
	Delete(c *gin.Context) error
}
