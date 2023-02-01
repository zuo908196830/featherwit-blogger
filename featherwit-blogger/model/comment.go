package model

import "time"

type Comment struct {
	ID           int       `json:"id" xorm:"int not null pk id autoincr"`
	CreateAt     time.Time `json:"createAt" xorm:"created"`
	UpdateAt     time.Time `json:"updateAt" xorm:"updated"`
	Username     string    `json:"username" xorm:"varchar(25)"` //评论者
	BlogId       int       `json:"blogId" xorm:"int"`
	ParentId     int       `json:"parentId" xorm:"int"` //如果是二级评论的话，所属一级评论id，一级评论时为负数
	ReplyId      int       `json:"replyId" xorm:"int"`  //如果是回复某二级评论的评论，二级评论id，否则为负数
	Content      string    `json:"Content" xorm:"text"`
	LikeCount    int64     `json:"likeCount" xorm:"bigint"`
	CommentCount int64     `json:"commentCount" xorm:"bigint"`
}
