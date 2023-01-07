package model

import "time"

type User struct {
	ID        int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt  time.Time `json:"createAt" xorm:"created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"updated"`
	Username  string    `json:"userName" xorm:"varchar(25) 'username' index"`
	Password  string    `json:"password" xorm:"varchar(45) 'password'"`
	Role      int       `json:"role" xorm:"int"`
	Nickname  string    `json:"nickname" xorm:"varchar(45)"`
	Telephone string    `json:"telephone" xorm:"varchar(20)"`
	Mail      string    `json:"mail" xorm:"varchar(45)"`
	Profile   string    `json:"profile" xorm:"varchar(400)"` //简介
	Headshot  string    `json:"headshot" xorm:"varchar(50)"`
}

func (u *User) TableName() string {
	return "users"
}
