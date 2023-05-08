package model

import "time"

type StarList struct {
	Username string    `json:"username" xorm:"varchar(25) not null index"`
	BlogId   int64     `json:"blogId" xorm:"bigint not null"`
	CreateAt time.Time `json:"createAt" xorm:"created"`
	UpdateAt time.Time `json:"updateAt" xorm:"updated"`
	Charset  string    `xorm:"'ENGINE=InnoDB CHARSET=utf8'"`
}

func (s *StarList) TableName() string {
	return "star_list"
}
