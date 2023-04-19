package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/response"
	"log"
	"xorm.io/xorm"
)

type TagService struct{}

var TagServiceApp = new(TagService)

func (ts *TagService) AddTag(tags *[]*model.Tag, s *xorm.Session) (int64, error) {
	s = CommonServiceApp.SetSession(s)
	n, err := s.InsertMulti(tags)
	if err != nil {
		log.Printf("insert tag error:%v", err)
		return 0, err
	}
	return n, nil
}

func (ts *TagService) AddTagBlog(blogID int64, tags *[]int64, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	for _, tag := range *tags {
		tagBlog := &model.TagBlog{
			BlogId: blogID,
			TagId:  tag,
		}
		exist, err := s.Table("tags").Exist(&model.Tag{ID: tag})
		if err != nil {
			log.Printf("select tag error:%v", err)
			return err
		} else if !exist {
			log.Printf("tag not exist")
			return errors.NewError(errors.ResourceNotExist, "tag not exist tagId:"+string(tag))
		}
		has, err := s.Exist(tagBlog)
		if err != nil {
			log.Printf("select tag_blog error:%v", err)
			return err
		}
		if !has {
			if _, err := s.Insert(tagBlog); err != nil {
				log.Printf("insert tag_blog error:%v", err)
				return err
			}
		}
	}
	return nil
}

func (ts *TagService) SearchTag(s *xorm.Session) (*response.SearchTagResponse, error) {
	s = CommonServiceApp.SetSession(s)
	tags := make([]*model.Tag, 0)
	err := s.Desc("search_count").Find(&tags)
	if err != nil {
		return nil, err
	}
	return &response.SearchTagResponse{Tags: tags}, nil
}

func (ts *TagService) DeleteTagBlogByBlogId(blogId int64, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	var tagBlog model.TagBlog
	num, err := s.Where("blog_id = ?", blogId).Count(&tagBlog)
	if err != nil {
		log.Printf("select tagBlog error:%v", err)
		return false, err
	}
	deleteNum, err := s.Where("blog_id = ?", blogId).Delete(&tagBlog)
	if err != nil {
		log.Printf("delete tagBlog error:%v", err)
		return false, err
	}
	if num != deleteNum {
		log.Printf("delete tagBlog error:%v", err)
		return false, nil
	}
	return true, nil
}

func (ts *TagService) AddSearchCount(tagId int64, s *xorm.Session) {
	s = CommonServiceApp.SetSession(s)
	tag := &model.Tag{ID: tagId}
	_, err := s.Cols("search_count").Get(tag)
	if err != nil {
		log.Printf("select tag error:%v", err)
		return
	}
	tag.SearchCount++
	s.Where("id = ?", tag.ID).Update(tag)
}
