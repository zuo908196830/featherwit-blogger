package global

import (
	"featherwit-blogger/model"
	"github.com/garyburd/redigo/redis"
	"log"
	"xorm.io/xorm"
)

var (
	DbEngine  *xorm.Engine
	RedisConn redis.Conn
)

func NewDbEngine() {
	db, err := xorm.NewEngine("mysql", "root:zuo123456789ke@tcp(127.0.0.1:3306)/the_fool_blogger")
	if err != nil {
		log.Fatalf("init db failed, err:%v", err)
	} else {
		DbEngine = db
	}
}

func InitDbEngine() {
	err := DbEngine.Sync2(
		new(model.User),
		new(model.Blob),
	)
	if err != nil {
		log.Fatalf("create table error:%v", err)
	}
}

func NewRedisEngine() {
	password := redis.DialPassword("zuo123456789ke")
	database := redis.DialDatabase(0)
	conn, err := redis.Dial("tcp", "121.89.220.131:6379", database, password)
	if err != nil {
		log.Fatalf("redis conn error:%v", err)
	}
	RedisConn = conn
}

func InitGlobal() {
	NewDbEngine()
	InitDbEngine()
	NewRedisEngine()
}
