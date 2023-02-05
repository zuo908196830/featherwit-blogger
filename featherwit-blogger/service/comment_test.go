package service

import (
	"featherwit-blogger/global"
	"log"
	"testing"
)

func TestSearchCommentByParentId(t *testing.T) {
	global.InitGlobal()
	idList, err := CommentServiceApp.SearchCommentByParentId(1, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range idList {
		log.Printf("%d ", val.ID)
	}
}

func TestDeleteCommentById(t *testing.T) {
	global.InitGlobal()
	idList := []int64{5, 7}
	err := CommentServiceApp.DeleteCommentById(idList, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
