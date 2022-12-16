package global

import (
	"log"
	"the-fool-blogger/model"
	"xorm.io/xorm"
)

var (
	DbEngine *xorm.Engine
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
	err := DbEngine.Sync2(new(model.User))
	if err != nil {
		log.Fatalf("create table error:%v", err)
	}
}
