package global

import (
	"log"
	"xorm.io/xorm"
)

var (
	DbEngine *xorm.Engine
)

func NewDbEngine() {
	db, err := xorm.NewEngine("mysql", "root:zuo123456789ke@tcp(127.0.0.1:3306)/redis_sgg")
	if err != nil {
		log.Fatalf("init db failed, err:%v", err)
	} else {
		DbEngine = db
	}
}
