package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"log"
)

type CommentService struct{}

var CommentServiceApp = new(CommentService)

func (cs *CommentService) AddComment(comment *model.Comment) error {
	_, err := global.DbEngine.Insert(comment)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (cs *CommentService) ContentCountPlus1(id int) (bool, error) {
	comment := &model.Comment{}
	ok, err := global.DbEngine.Cols("comment_count").Where("id = ?", id).Get(comment)
	if err != nil {
		log.Printf("get comment error:%v", err)
		return false, err
	} else if !ok {
		log.Printf("comment is not exist")
		return false, nil
	}
	mp := make(map[string]int64)
	mp["comment_count"] = comment.CommentCount + 1
	i, err := global.DbEngine.Table(comment).Where("id = ?", id).Update(mp)
	if err != nil {
		log.Printf("update comment error:%v", err)
		return false, err
	} else if i == 0 {
		log.Printf("update comment is 0")
		return false, nil
	}
	return true, nil
}

func (cs *CommentService) GetCommentById(id int) (*model.Comment, error) {
	comment := &model.Comment{}
	has, err := global.DbEngine.Where("id = ?", id).Get(comment)
	if err != nil {
		log.Printf("select comment error:%v", err)
		return nil, err
	} else if !has {
		log.Printf("comment is not exist id:%d", id)
		return nil, nil
	}
	return comment, nil
}

func (cs *CommentService) CommentExist(id int) (bool, error) {
	has, err := global.DbEngine.Exist(&model.Comment{ID: id})
	if err != nil {
		log.Printf("select comment error:%v", err)
		return false, err
	}
	return has, nil
}
