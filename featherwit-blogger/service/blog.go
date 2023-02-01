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
	err := global.DbEngine.Cols("id", "username", "create_at", "update_at", "title", "views", "comment_count", "like_count").Limit(limit, offset).Find(&blogs)
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

func (b *BlogService) BlogExist(id int) (bool, error) {
	exist, err := global.DbEngine.Exist(&model.Blob{ID: id})
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

func (b *BlogService) BlogCount() (int64, error) {
	n, err := global.DbEngine.Count(new(model.Blob))
	if err != nil {
		log.Printf("get blog error:%v", err)
		return 0, err
	}
	return n, nil
}

func (b *BlogService) ContentCountPlus1(blogId int) (bool, error) {
	blog := &model.Blob{}
	ok, err := global.DbEngine.Cols("comment_count").Where("id = ?", blogId).Get(blog)
	if err != nil {
		log.Printf("get blog error:%v", err)
		return false, err
	} else if !ok {
		log.Printf("blog is not exist")
		return false, nil
	}
	mp := make(map[string]int64)
	mp["comment_count"] = blog.CommentCount + 1
	i, err := global.DbEngine.Table(blog).Where("id = ?", blogId).Update(mp)
	if err != nil {
		log.Printf("update blog error:%v", err)
		return false, err
	} else if i == 0 {
		log.Printf("update blog is 0")
		return false, nil
	}
	return true, nil
}
