package model

import "time"

type User struct {
	ID       string    `json:"id" xorm:"int not null pk id autoincr"`
	Username string    `json:"userName" xorm:"varchar(45) 'username'"`
	Password string    `json:"password" xorm:"varchar(45) 'password'"`
	CreateAt time.Time `json:"createAt" xorm:"created"`
	UpdateAt time.Time `json:"updateAt" xorm:"updated"`
	Role     int       `json:"role" xorm:"int"`
	Nickname string    `json:"nickname" xorm:"varchar(50)"`
}

func (u *User) TableName() string {
	return "users"
}
