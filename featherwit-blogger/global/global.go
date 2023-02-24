package global

import (
	"featherwit-blogger/model"
	"log"

	"github.com/garyburd/redigo/redis"
	"xorm.io/xorm"
)

var (
	DbEngine      *xorm.Engine
	RedisConnPool *redis.Pool
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
	DbEngine.ShowSQL(true)
	err := DbEngine.Sync2(
		new(model.User),
		new(model.Blog),
		new(model.Comment),
		new(model.Tag),
		new(model.TagBlog),
	)
	if err != nil {
		log.Fatalf("create table error:%v", err)
	}
}

func NewRedisEngine() {
	password := redis.DialPassword("zuo123456789ke")
	database := redis.DialDatabase(0)
	pool := &redis.Pool{
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "121.89.220.131:6379", database, password)
		},
	}
	RedisConnPool = pool
}

func InitGlobal() {
	NewDbEngine()
	InitDbEngine()
	NewRedisEngine()
}
