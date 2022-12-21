package service

import (
	"encoding/json"
	"featherwit-blogger/global"
	"log"
)

type CommonService struct{}

var CommonServiceApp = new(CommonService)

func (c *CommonService) RedisSet(key string, val interface{}) error {
	_, ok1 := val.(string)
	_, ok2 := val.(int)
	if ok1 || ok2 {
		_, err := global.RedisConn.Do("Set", key, val)
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
		_, err = global.RedisConn.Do("Set", key, v)
		if err != nil {
			log.Printf("redis set error:%v", err)
			return err
		}
		return nil
	}
}

func (c *CommonService) RedisSetTime(key string, t int) error {
	_, err := global.RedisConn.Do("expire", key, t)
	if err != nil {
		log.Printf("set redis expiration time error:%v", err)
		return err
	}
	return nil
}

func (c *CommonService) RedisGet(key string) (interface{}, error) {
	val, err := global.RedisConn.Do("Get", key)
	if err != nil {
		log.Printf("redis get error:%v", err)
		return nil, err
	}
	return val, nil
}

func (c *CommonService) RedisDelete(key string) error {
	_, err := global.RedisConn.Do("del", key)
	if err != nil {
		log.Printf("redis delete error: %v", err)
		return err
	}
	return nil
}
