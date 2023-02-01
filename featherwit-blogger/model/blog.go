package model

import "time"

type Blob struct {
	ID           int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt     time.Time `json:"createAt" xorm:"created"`
	UpdateAt     time.Time `json:"updateAt" xorm:"updated"`
	Username     string    `json:"username" xorm:"varchar(25) index"` //作者id
	Title        string    `json:"title" xorm:"text"`
	Content      string    `json:"content" xrom:"longtext"`
	Views        int64     `json:"views" xorm:"bigint not null default 0"` //浏览量
	CommentCount int64     `json:"commentCount" xorm:"bigint not null default 0"`
	LikeCount    int64     `json:"likeCount" xorm:"bigint not null default 0"`
	Cover        string    `json:"cover" xorm:"varchar(200)"` //封面
	Profile      string    `json:"profile" xorm:"text"`       //文章简介
}

func (b *Blob) TableName() string {
	return "blog"
}
