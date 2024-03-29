package model

import "time"

type Comment struct {
	ID           int64     `json:"id" xorm:"bigint not null pk id autoincr"`
	CreateAt     time.Time `json:"createAt" xorm:"created"`
	UpdateAt     time.Time `json:"updateAt" xorm:"updated"`
	Username     string    `json:"username" xorm:"varchar(25) index"` //评论者
	BlogId       int64     `json:"blogId" xorm:"bigint index"`
	ParentId     int64     `json:"parentId" xorm:"bigint default -1"` //如果是二级评论的话，所属一级评论id，一级评论时为负数
	ReplyId      int64     `json:"replyId" xorm:"bigint default -1"`
	Content      string    `json:"content" xorm:"text"`
	LikeCount    int64     `json:"likeCount" xorm:"bigint default 0"`
	CommentCount int64     `json:"commentCount" xorm:"bigint default 0"`
}

func (c *Comment) TableName() string {
	return "comment"
}
