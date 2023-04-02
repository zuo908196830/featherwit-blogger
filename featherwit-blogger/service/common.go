package service

import (
	"crypto/rand"
	"encoding/json"
	"featherwit-blogger/global"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

func (cs *CommonService) GetFileName() (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	s := fmt.Sprintf("%x", b)
	return s, err
}

func (cs *CommonService) UploadImg(img *multipart.File, headers *multipart.FileHeader) (string, error) {
	Endpoint := "https://" + global.GlobalConfig.AccessKey.Endpoint
	AccessKeyId := global.GlobalConfig.AccessKey.AccessKeyId
	AccessKeySecret := global.GlobalConfig.AccessKey.AccessKeySecret
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		log.Printf("get oss client error :%v", err)
		return "", err
	}
	bucket, err := client.Bucket(global.GlobalConfig.AccessKey.ImgBucket)
	if err != nil {
		log.Printf("get oss bucket error :%v", err)
		return "", err
	}

	// todo 识别文件后缀，生成一个随机、唯一的文件名，添加上原后缀
	imgName := ""
	ok := true
	for ok {
		parts := strings.Split(headers.Filename, ".")
		last := parts[len(parts)-1]
		s, _ := cs.GetFileName()
		imgName = "img/" + s + "." + last
		ok, err = bucket.IsObjectExist(imgName)
		if err != nil {
			log.Printf("search oss img error :%v", err)
			return "", err
		}
	}
	err = bucket.PutObject(imgName, *img)
	if err != nil {
		log.Printf("save oss img error :%v", err)
		return "", err
	}
	imgUrl := "https://" + global.GlobalConfig.AccessKey.ImgBucket + "." + global.GlobalConfig.AccessKey.Endpoint + "/" + imgName
	return imgUrl, nil
}
