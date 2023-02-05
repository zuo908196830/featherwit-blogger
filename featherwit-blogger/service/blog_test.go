package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBlogService_AddBlog(t *testing.T) {
	global.InitGlobal()
	blog := &model.Blob{
		Username:     "123",
		Title:        "123",
		Content:      "123",
		Views:        0,
		CommentCount: 0,
		LikeCount:    0,
	}
	err := BlogServiceApp.AddBlog(blog, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", blog)
}

func TestBlogService_SearchBlog(t *testing.T) {
	global.InitGlobal()
	blogs, err := BlogServiceApp.SearchBlog(2, 1, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, blog := range blogs {
		log.Printf("%v", blog)
	}
}

func TestBlogService_GetBlogById(t *testing.T) {
	global.InitGlobal()
	blob, err := BlogServiceApp.GetBlogById(1, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", blob)
}

func TestBlogService_BlogExist(t *testing.T) {
	global.InitGlobal()
	exist, err := BlogServiceApp.BlogExist(1, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println(exist)
}

func TestBlogService_ContentCountPlus1(t *testing.T) {
	global.InitGlobal()
	b, err := BlogServiceApp.UpdateCommentCount(1, 1, nil)
	if err != nil {
		log.Fatalf("%v", err)
	} else if !b {
		log.Fatal()
	}
	log.Printf("ok")
}
