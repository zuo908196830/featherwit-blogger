package model

import "time"

type TagBlog struct {
	TagId1   string    `json:"tagId1" xorm:"varchar(6)"`
	TagId2   string    `json:"tagId2" xorm:"varchar(6)"`
	TagId3   string    `json:"tagId3" xorm:"varchar(6)"`
	Height   int       `json:"height" xorm:"int"`
	BlogId   int64     `json:"blogId" xorm:"bigint index"`
	CreateAt time.Time `json:"createAt" xorm:"created"`
	UpdateAt time.Time `json:"updateAt" xorm:"updated"`
}
