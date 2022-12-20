package model

import "time"

type BaseModel struct {
	ID       int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt time.Time `json:"createAt" xorm:"created"`
	UpdateAt time.Time `json:"updateAt" xorm:"updated"`
}
