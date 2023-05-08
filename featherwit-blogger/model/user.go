package model

import "time"

type User struct {
	CreateAt  time.Time `json:"createAt" xorm:"created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"updated"`
	Username  string    `json:"userName" xorm:"varchar(25) 'username' index unique"`
	Password  string    `json:"password" xorm:"varchar(45) 'password'"`
	Role      int       `json:"role" xorm:"int"`
	Nickname  string    `json:"nickname" xorm:"varchar(45)"`
	Telephone string    `json:"telephone" xorm:"varchar(20)"`
	Mail      string    `json:"mail" xorm:"varchar(45)"`
	Profile   string    `json:"profile" xorm:"varchar(400)"` //简介
	Headshot  string    `json:"headshot" xorm:"varchar(200)"`
	FansCount int64     `json:"fansCount" xorm:"bigint not null default 0"`
	Charset   string    `xorm:"'ENGINE=InnoDB CHARSET=utf8'"`
}

func (u *User) TableName() string {
	return "users"
}
