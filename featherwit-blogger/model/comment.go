package model

import "time"

type Comment struct {
	ID        int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt  time.Time `json:"createAt" xorm:"created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"updated"`
	Username  string    `json:"username" xorm:"varchar(25)"` //评论者
	Type      int       `json:"type" xorm:"int"`             //0为基础评论 1为评论基础评论的评论 2为回复1类评论的评论
	ParentId  int       `json:"parentId" xorm:"int"`
	Content   string    `json:"Content" xorm:"text"`
	LikeCount int64     `json:"likeCount" xorm:"bigint"`
}
