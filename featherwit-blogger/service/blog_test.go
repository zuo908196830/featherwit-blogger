package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestBlogService_AddBlog(t *testing.T) {
	global.InitGlobal()
	blog := &model.Blob{
		Username:    "123",
		Title:       "123",
		Content:     "123",
		Views:       0,
		CommonCount: 0,
		LikeCount:   0,
	}
	err := BlogServiceApp.AddBlog(blog)
	if err != nil {
		log.Fatalf("asd")
	}
	log.Printf("%v", blog)
}
