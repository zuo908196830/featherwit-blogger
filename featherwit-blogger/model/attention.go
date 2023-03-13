package model

import "time"

type Attention struct {
	Username  string    `json:"username" xorm:"varchar(25) index"`
	AUsername string    `json:"aUsername" xorm:"varchar(25)"`
	CreateAt  time.Time `json:"createAt" xorm:"created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"updated"`
}

func (a *Attention) TableName() string {
	return "attention"
}
