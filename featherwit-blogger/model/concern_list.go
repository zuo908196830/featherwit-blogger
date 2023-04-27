package model

import "time"

type ConcernList struct {
	Username        string    `json:"username" xorm:"varchar(25) not null index"`
	ConcernUsername string    `json:"concernUsername" xorm:"varchar(25) not null"`
	CreateAt        time.Time `json:"createAt" xorm:"created"`
	UpdateAt        time.Time `json:"updateAt" xorm:"updated"`
}
