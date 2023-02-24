package service

import (
	"featherwit-blogger/model"
	"log"

	"xorm.io/xorm"
)

type BlogService struct{}

var BlogServiceApp = new(BlogService)

func (b *BlogService) AddBlog(blog *model.Blob, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Insert(blog)
	if err != nil {
		log.Printf("add blog error: %v", err)
		return err
	}
	return nil
}

func (b *BlogService) SearchBlog(limit int, offset int, s *xorm.Session) ([]*model.Blob, error) {
	s = CommonServiceApp.SetSession(s)
	blogs := make([]*model.Blob, 0)
	err := s.Cols("id", "username", "create_at", "update_at", "title", "views", "comment_count", "like_count").Limit(limit, offset).Find(&blogs)
	if err != nil {
		log.Printf("search blog error: %v", err)
		return nil, err
	}
	return blogs, nil
}

func (b *BlogService) GetBlogById(id int64, s *xorm.Session) (*model.Blob, error) {
	s = CommonServiceApp.SetSession(s)
	blog := &model.Blob{ID: id}
	ok, err := s.Get(blog)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}
	return blog, nil
}

func (b *BlogService) BlogExist(id int64, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	exist, err := s.Exist(&model.Blob{ID: id})
	if err != nil {
		return false, err
	}
	return exist, err
}

func (b *BlogService) UpdateBlog(blog *model.Blob, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Where("id = ?", blog.ID).Update(blog)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogService) BlogCount(s *xorm.Session) (int64, error) {
	s = CommonServiceApp.SetSession(s)
	n, err := s.Count(new(model.Blob))
	if err != nil {
		log.Printf("get blog error:%v", err)
		return 0, err
	}
	return n, nil
}

func (b *BlogService) UpdateCommentCount(blogId int64, num int, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	blog := &model.Blob{}
	ok, err := s.Cols("comment_count").Where("id = ?", blogId).Get(blog)
	if err != nil {
		log.Printf("get blog error:%v", err)
		return false, err
	} else if !ok {
		log.Printf("blog is not exist")
		return false, nil
	}
	mp := make(map[string]int64)
	mp["comment_count"] = blog.CommentCount + int64(num)
	i, err := s.Table(blog).Where("id = ?", blogId).Update(mp)
	if err != nil {
		log.Printf("update blog error:%v", err)
		return false, err
	} else if i == 0 {
		log.Printf("update blog is 0")
		return false, nil
	}
	return true, nil
}
