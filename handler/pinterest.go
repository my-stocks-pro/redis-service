package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"net/http"
	"github.com/pkg/errors"
)

type Pinterest interface {
	Get(c *gin.Context) error
	Post(c *gin.Context) error
	Put(c *gin.Context) error
	Delete(c *gin.Context) error
}

type TypePinterest struct {
	config infrastructure.Config
	logger infrastructure.Logger
	redis  infrastructure.Redis
}

func NewPinterest(c infrastructure.Config, l infrastructure.Logger, r infrastructure.Redis) TypePinterest {
	return TypePinterest{
		config: c,
		logger: l,
		redis:  r,
	}
}

//func (p TypePinterest) Handle(c *gin.Context) {
//	db, body, err := p.getParams(c)
//	if err != nil {
//		p.logger.ContextError(c, http.StatusInternalServerError, err)
//		return
//	}
//
//	switch c.Request.Method {
//	case http.MethodGet:
//		if err := p.get(c, body.Key, db); err != nil {
//			p.logger.ContextError(c, http.StatusInternalServerError, err)
//			return
//		}
//	case http.MethodPost:
//		if err := p.post(c, body.Key, body.Val, db); err != nil {
//			p.logger.ContextError(c, http.StatusInternalServerError, err)
//			return
//		}
//	case http.MethodPut:
//		if err := p.put(c, body.Key, body.Val, db); err != nil {
//			p.logger.ContextError(c, http.StatusInternalServerError, err)
//			return
//		}
//	case http.MethodDelete:
//		if err := p.delete(c, body.Key, db); err != nil {
//			p.logger.ContextError(c, http.StatusInternalServerError, err)
//			return
//		}
//	default:
//		p.logger.ContextError(c, http.StatusMethodNotAllowed, errors.New("Method Not Allowed"))
//		return
//	}
//
//	p.logger.ContextSuccess(c, http.StatusOK)
//}

func (p TypePinterest) Get(c *gin.Context) error {

	db, body, err := getPara

	lLen, err := p.redis.LLen(key, db)
	if err != nil {
		return err
	}

	if lLen == 0 {
		return errors.Errorf("RedisDB: %s / KEY: %s / LLen == %d", db, key, lLen)
	}

	val, err := p.redis.LPop(key, db)
	if err != nil {
		return err
	}

	_, err = c.Writer.Write(val)
	if err != nil {
		return err
	}

	return nil
}

func (p TypePinterest) Post(c *gin.Context) error {
	if err := p.redis.LPush(key, val, db); err != nil {
		return err
	}

	return nil
}

func (p TypePinterest) Put(c *gin.Context) error {
	if err := p.redis.RPush(key, val, db); err != nil {
		return err
	}

	return nil
}

func (p TypePinterest) Delete(c *gin.Context) error {
	err := p.redis.Delete(key, db)
	if err != nil {
		return err
	}

	return nil
}
