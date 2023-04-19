package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/request"
	"log"

	"xorm.io/xorm"
)

type BlogService struct{}

var BlogServiceApp = new(BlogService)

func (b *BlogService) AddBlog(blog *model.Blog, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Insert(blog)
	if err != nil {
		log.Printf("add blog error: %v", err)
		return err
	}
	return nil
}

func (b *BlogService) SearchBlog(param *request.SearchBlogRequest, limit int, offset int, s *xorm.Session) ([]*model.Blog, error) {
	s = CommonServiceApp.SetSession(s)
	blogs := make([]*model.Blog, 0)
	s = s.Table("blog").Alias("b")
	s = s.Cols("id", "username", "b.create_at as create_at", "b.update_at as update_at", "title", "views", "comment_count", "like_count", "profile", "cover").Limit(limit, offset).Where("1=1")
	if param.Name != "" {
		s = s.And("title like ?", "%"+param.Name+"%").Or("profile like ?", "%"+param.Name+"%")
	}
	if param.TagId != 0 {
		s = s.Join("INNER", "tag_blog", "id = blog_id").And("tag_id = ?", param.TagId)
	}
	err := s.Find(&blogs)
	if err != nil {
		log.Printf("search blog error: %v", err)
		return nil, err
	}
	return blogs, nil
}

func (b *BlogService) GetBlogById(id int64, s *xorm.Session) (*model.Blog, error) {
	s = CommonServiceApp.SetSession(s)
	blog := &model.Blog{ID: id}
	ok, err := s.ForUpdate().Get(blog)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}
	_, err = s.Where("id = ?", id).Cols("views").Update(&model.Blog{Views: blog.Views + 1})
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (b *BlogService) BlogExist(id int64, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	exist, err := s.Exist(&model.Blog{ID: id})
	if err != nil {
		return false, err
	}
	return exist, err
}

func (b *BlogService) UpdateBlog(blog *model.Blog, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Where("id = ?", blog.ID).Update(blog)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlogService) BlogCount(s *xorm.Session) (int64, error) {
	s = CommonServiceApp.SetSession(s)
	n, err := s.Count(new(model.Blog))
	if err != nil {
		log.Printf("get blog error:%v", err)
		return 0, err
	}
	return n, nil
}

func (b *BlogService) UpdateCommentCount(blogId int64, num int, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	blog := &model.Blog{}
	ok, err := s.Cols("comment_count").Where("id = ?", blogId).ForUpdate().Get(blog)
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

func (b *BlogService) DeleteBlogById(id int64, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	blog := &model.Blog{
		ID: id,
	}
	_, err := s.Delete(blog)
	if err != nil {
		log.Printf("delete blog error:%v", err)
		return err
	}
	return nil
}
