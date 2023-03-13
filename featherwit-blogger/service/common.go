package service

import (
	"encoding/json"
	"featherwit-blogger/global"
	"log"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type CommonService struct{}

var CommonServiceApp = new(CommonService)

func (c *CommonService) RedisSet(key string, val interface{}) error {
	RedisConn := global.RedisConnPool.Get()
	defer RedisConn.Close()
	_, ok1 := val.(string)
	_, ok2 := val.(int)
	if ok1 || ok2 {
		_, err := RedisConn.Do("Set", key, val)
		if err != nil {
			log.Printf("redis set error:%v", err)
			return err
		}
		return nil
	} else {
		b, err := json.Marshal(val)
		if err != nil {
			log.Printf("json error: %v", err)
			return err
		}
		v := string(b)
		_, err = RedisConn.Do("Set", key, v)
		if err != nil {
			log.Printf("redis set error:%v", err)
			return err
		}
		return nil
	}
}

func (c *CommonService) RedisSetTime(key string, t int) error {
	RedisConn := global.RedisConnPool.Get()
	defer RedisConn.Close()
	_, err := RedisConn.Do("expire", key, t)
	if err != nil {
		log.Printf("set redis expiration time error:%v", err)
		return err
	}
	return nil
}

func (c *CommonService) RedisGet(key string) (interface{}, error) {
	RedisConn := global.RedisConnPool.Get()
	defer RedisConn.Close()
	val, err := RedisConn.Do("Get", key)
	if err != nil {
		log.Printf("redis get error:%v", err)
		return nil, err
	}
	return val, nil
}

func (c *CommonService) RedisDelete(key string) error {
	RedisConn := global.RedisConnPool.Get()
	defer RedisConn.Close()
	_, err := RedisConn.Do("del", key)
	if err != nil {
		log.Printf("redis delete error: %v", err)
		return err
	}
	return nil
}

func (cs *CommonService) SetSession(s *xorm.Session) *xorm.Session {
	if s == nil {
		return global.DbEngine.NewSession()
	}
	return s
}

func (cs *CommonService) GetUsername(c *gin.Context) string {
	val, _ := c.Get("User-Info")
	tkmp := val.(map[string]interface{})
	username := tkmp["username"].(string)
	return username
}
