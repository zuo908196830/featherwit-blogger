package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/response"
	"log"

	"xorm.io/xorm"
)

type CommentService struct{}

var CommentServiceApp = new(CommentService)

func (cs *CommentService) AddComment(comment *model.Comment, s *xorm.Session) (*model.Comment, error) {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Insert(comment)
	if err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}

func (cs *CommentService) UpdateCommentCount(id int64, num int, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	comment := &model.Comment{}
	ok, err := s.Cols("comment_count").Where("id = ?", id).ForUpdate().Get(comment)
	if err != nil {
		log.Printf("get comment error:%v", err)
		return false, err
	} else if !ok {
		log.Printf("comment is not exist")
		return false, nil
	}
	mp := make(map[string]interface{})
	mp["comment_count"] = comment.CommentCount + int64(num)
	i, err := s.Table(comment).Where("id = ?", id).Update(mp)
	if err != nil {
		log.Printf("update comment error:%v", err)
		return false, err
	} else if i == 0 {
		log.Printf("update comment is 0")
		return false, nil
	}
	return true, nil
}

func (cs *CommentService) GetCommentById(id int64, s *xorm.Session) (*model.Comment, error) {
	s = CommonServiceApp.SetSession(s)
	comment := &model.Comment{}
	has, err := s.Where("id = ?", id).Get(comment)
	if err != nil {
		log.Printf("select comment error:%v", err)
		return nil, err
	} else if !has {
		log.Printf("comment is not exist id:%d", id)
		return nil, nil
	}
	return comment, nil
}

func (cs *CommentService) CommentExist(id int64, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	has, err := s.Exist(&model.Comment{ID: id})
	if err != nil {
		log.Printf("select comment error:%v", err)
		return false, err
	}
	return has, nil
}

func (cs *CommentService) SearchCommentByParentId(parentId []int64, s *xorm.Session) ([]*model.Comment, error) {
	s = CommonServiceApp.SetSession(s)
	commentIdList := make([]*model.Comment, 0)
	err := s.Cols("id").In("parent_id = ?", parentId).Find(&commentIdList)
	if err != nil {
		log.Printf("search id from comments error:%v", err)
		return nil, err
	}
	return commentIdList, err
}

func (cs *CommentService) DeleteCommentById(idList []int64, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	comment := &model.Comment{}
	_, err := s.In("id", idList).Delete(comment)
	if err != nil {
		log.Printf("delete comment error:%v", err)
		return err
	}
	return nil
}

func (cs *CommentService) DeleteCommentByBlogId(blogId int64, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	var comment model.Comment
	num, err := s.Where("blog_id = ?", blogId).Count(&comment)
	if err != nil {
		log.Printf("select comment count error:%v", err)
		return false, err
	}
	deleteNum, err := s.Where("blog_id = ?", blogId).Delete(comment)
	if err != nil {
		log.Printf("delete comment error:%v", err)
		return false, err
	}
	if num != deleteNum {
		log.Printf("select num != delete num")
		return false, nil
	}
	return true, nil
}

func (cs *CommentService) GetCommentByBlogId(blogId int64, limit int, offset int, s *xorm.Session) ([]*model.Comment, error) {
	s = CommonServiceApp.SetSession(s)
	comments := make([]*model.Comment, 0)
	err := s.Where("blog_id = ?", blogId).Desc("create_at").And("parent_id < 0").Limit(limit, offset).Find(&comments)
	if err != nil {
		log.Printf("select comment by blog id error:%v", err)
		return nil, err
	}
	return comments, nil
}

func (cs *CommentService) GetCommentByParentID(parentId int64, s *xorm.Session) ([]*response.ChildrenComment, error) {
	s = CommonServiceApp.SetSession(s)
	comments := make([]*model.Comment, 0)
	err := s.Where("parent_id = ?", parentId).Desc("create_at").Find(&comments)
	if err != nil {
		log.Printf("select comment by parentId error:%v", err)
		return nil, err
	}
	res := make([]*response.ChildrenComment, len(comments))
	for i := 0; i < len(comments); i++ {
		res[i] = new(response.ChildrenComment)
		res[i].Comment = comments[i]
		nickname, headshot, err := UserServiceApp.GetNicknameAndCover(comments[i].Username, s)
		if err != nil {
			return nil, err
		}
		res[i].User = &response.UserShow{
			Nickname: nickname,
			Headshot: headshot,
		}
	}
	return res, nil
}

func (cs *CommentService) GetCount(blogId int64, s *xorm.Session) (int, error) {
	s = CommonServiceApp.SetSession(s)
	n, err := s.Where("blog_id = ?", blogId).Table("comment").Count()
	if err != nil {
		log.Printf("select comment error:%v", err)
		return 0, err
	}
	return int(n), nil
}
