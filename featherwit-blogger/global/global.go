package global

import (
	"encoding/json"
	"featherwit-blogger/model"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"xorm.io/xorm"
)

var (
	DbEngine      *xorm.Engine
	RedisConnPool *redis.Pool
	GlobalConfig  *Config
)

func InitConfig() {
	data, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &GlobalConfig)
	if err != nil {
		panic(err)
	}
	log.Printf("Name: %s, Age: %d\n", GlobalConfig.MySQL.Ip, GlobalConfig.Redis.Database)
}

func NewDbEngine() {
	url := GlobalConfig.MySQL.Username + ":" + GlobalConfig.MySQL.Password + "@tcp(" + GlobalConfig.MySQL.Ip + ":" + fmt.Sprintf("%d)/", GlobalConfig.MySQL.Port) + GlobalConfig.MySQL.Database
	db, err := xorm.NewEngine("mysql", url)
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
		new(model.Attention),
	)
	if err != nil {
		log.Fatalf("create table error:%v", err)
	}
}

func NewRedisEngine() {
	password := redis.DialPassword(GlobalConfig.Redis.Password)
	database := redis.DialDatabase(GlobalConfig.Redis.Database)
	pool := &redis.Pool{
		MaxIdle: GlobalConfig.Redis.MaxIdle, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   GlobalConfig.Redis.MaxActive,                  //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: time.Duration(GlobalConfig.Redis.IdleTimeout), //连接关闭时间 300秒 （300秒不使用自动关闭）

		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", GlobalConfig.Redis.Ip, GlobalConfig.Redis.Port), database, password)
		},
	}
	RedisConnPool = pool
}

func InitGlobal() {
	InitConfig()
	NewDbEngine()
	InitDbEngine()
	NewRedisEngine()
}
