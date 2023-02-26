package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"log"
	"math/rand"

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

func (ts *TagService) MutilAddTag(tagTree []*request.TagTree, id1 string, id2 string, height int, tags *[]*model.Tag) {
	if len(tagTree) == 0 {
		return
	}
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, tag := range tagTree {
		var id string
		for j := 0; j < 6; j++ {
			id += string(charset[rand.Intn(len(charset))])
		}
		if height == 1 {
			*tags = append(*tags, &model.Tag{Id1: id, Id2: "", Id3: "", Name: tag.Name, Height: 1})
			ts.MutilAddTag(tag.Children, id, "", 2, tags)
		} else if height == 2 {
			*tags = append(*tags, &model.Tag{Id1: id1, Id2: id, Id3: "", Name: tag.Name, Height: 2})
			ts.MutilAddTag(tag.Children, id1, id, 3, tags)
		} else {
			*tags = append(*tags, &model.Tag{Id1: id1, Id2: id2, Id3: id, Name: tag.Name, Height: 3})
		}
	}
}

func (ts *TagService) AddTagBlog(blogID int64, tags *[]*request.TagRequest, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	for _, tag := range *tags {
		tagBlog := &model.TagBlog{
			BlogId: blogID,
			TagId1: tag.ID1,
			TagId2: tag.ID2,
			TagId3: tag.ID3,
			Height: tag.Height,
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

func (ts *TagService) SearchTag(s *xorm.Session) (*[]*response.SearchTagResponse, error) {
	s = CommonServiceApp.SetSession(s)
	tagTree1 := make([]*model.Tag, 0)
	res := make([]*response.SearchTagResponse, 0)
	err := s.Where("height = ?", 1).Find(&tagTree1)
	if err != nil {
		log.Printf("search tag error:%v", err)
		return nil, err
	}
	for _, tag := range tagTree1 {
		res = append(res, &response.SearchTagResponse{
			Tag: tag,
		})
	}
	for i := 0; i < len(res); i++ {
		tagTree2 := make([]*model.Tag, 0)
		err := s.Where("height = ?", 2).And("id1 = ?", res[i].Tag.Id1).Find(&tagTree2)
		if err != nil {
			log.Printf("search tag error:%v", err)
			return nil, err
		}
		if len(tagTree2) != 0 {
			res[i].Children = make([]*response.SearchTagResponse, 0)
			for _, tag := range tagTree2 {
				res[i].Children = append(res[i].Children, &response.SearchTagResponse{
					Tag: tag,
				})
			}
		}
		for j := 0; j < len(res[i].Children); j++ {
			tagTree3 := make([]*model.Tag, 0)
			err := s.Where("height = ?", 3).And("id1 = ?", res[i].Children[j].Tag.Id1).And("id2 = ?", res[i].Children[j].Tag.Id2).Find(&tagTree3)
			if err != nil {
				log.Printf("search tag error:%v", err)
				return nil, err
			}
			if len(tagTree3) != 0 {
				res[i].Children[j].Children = make([]*response.SearchTagResponse, 0)
				for _, tag2 := range tagTree3 {
					res[i].Children[j].Children = append(res[i].Children[j].Children, &response.SearchTagResponse{
						Tag: tag2,
					})
				}
			}
		}
	}
	return &res, nil
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
