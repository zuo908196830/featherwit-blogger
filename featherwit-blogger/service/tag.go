package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/request"
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
			*tags = append(*tags, &model.Tag{ID1: id, ID2: "", ID3: "", Name: tag.Name, Height: 1})
			ts.MutilAddTag(tag.Children, id, "", 2, tags)
		} else if height == 2 {
			*tags = append(*tags, &model.Tag{ID1: id1, ID2: id, ID3: "", Name: tag.Name, Height: 2})
			ts.MutilAddTag(tag.Children, id1, id, 3, tags)
		} else {
			*tags = append(*tags, &model.Tag{ID1: id1, ID2: id2, ID3: id, Name: tag.Name, Height: 3})
		}
	}
}
