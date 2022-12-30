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

func (b *BlogService) SearchBlog(limit int, offset int) ([]*model.Blob, error) {
	blogs := make([]*model.Blob, 0)
	err := global.DbEngine.Cols("id", "username", "create_at", "update_at", "title", "views", "common_count", "like_count").Limit(limit, offset).Find(&blogs)
	if err != nil {
		log.Printf("search blog error: %v", err)
		return nil, err
	}
	return blogs, nil
}

func (b *BlogService) GetBlogById(id int) (*model.Blob, error) {
	blog := &model.Blob{ID: id}
	ok, err := global.DbEngine.Get(blog)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}
	return blog, nil
}

func (b *BlogService) BlogExist(id int, username string) (bool, error) {
	exist, err := global.DbEngine.Exist(&model.Blob{ID: id, Username: username})
	if err != nil {
		return false, err
	}
	return exist, err
}

func (b *BlogService) UpdateBlog(blog *model.Blob) error {
	_, err := global.DbEngine.Where("id = ?", blog.ID).Update(blog)
	if err != nil {
		return err
	}
	return nil
}
