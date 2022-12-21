package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"log"
)

type BlogService struct{}

var BlogServiceApp = new(BlogService)

func (b *BlogService) AddBlog(blog *model.Blob) error {
	_, err := global.DbEngine.Insert(blog)
	if err != nil {
		log.Printf("add blog error: %v", err)
		return err
	}
	return nil
}
