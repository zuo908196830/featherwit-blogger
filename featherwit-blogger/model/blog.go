package model

import "time"

type Blob struct {
	ID          int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt    time.Time `json:"createAt" xorm:"created"`
	UpdateAt    time.Time `json:"updateAt" xorm:"updated"`
	Username    string    `json:"username" xorm:"varchar(25) index"`
	Title       string    `json:"title" xorm:"text"`
	Content     string    `json:"content" xrom:"longtext"`
	Views       int64     `json:"views" xorm:"bigint"` //浏览量
	CommonCount int64     `json:"commonCount" xorm:"bigint"`
	LikeCount   int64     `json:"likeCount" xorm:"bigint"`
	Cover       string    `json:"cover" xorm:"varchar(200)"`
}

func (b *Blob) TableName() string {
	return "blog"
}
