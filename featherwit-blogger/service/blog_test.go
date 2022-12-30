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
		log.Fatalf("%v", err)
	}
	log.Printf("%v", blog)
}

func TestBlogService_SearchBlog(t *testing.T) {
	global.InitGlobal()
	blogs, err := BlogServiceApp.SearchBlog(2, 1)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, blog := range blogs {
		log.Printf("%v", blog)
	}
}

func TestBlogService_GetBlogById(t *testing.T) {
	global.InitGlobal()
	blob, err := BlogServiceApp.GetBlogById(1)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", blob)
}

func TestBlogService_BlogExist(t *testing.T) {
	global.InitGlobal()
	exist, err := BlogServiceApp.BlogExist(1, "xxx")
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println(exist)
}
