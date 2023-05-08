package model

import "time"

type TagBlog struct {
	TagId    int64     `json:"tagId" xorm:"bigint not null"`
	BlogId   int64     `json:"blogId" xorm:"bigint index not null"`
	CreateAt time.Time `json:"createAt" xorm:"created"`
	UpdateAt time.Time `json:"updateAt" xorm:"updated"`
	Charset  string    `xorm:"'ENGINE=InnoDB CHARSET=utf8'"`
}

func (tb *TagBlog) TableName() string {
	return "tag_blog"
}
