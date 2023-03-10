package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"log"

	"xorm.io/xorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) GetUserByUsername(username string, s *xorm.Session) (*model.User, error) {
	s = CommonServiceApp.SetSession(s)
	user := new(model.User)
	ok, err := s.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	} else {
		return user, nil
	}
}

func (u *UserService) AddUser(user *model.User, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	_, err := s.Insert(user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserService) LoginStatus(user map[string]interface{}, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	username, ok := user["username"]
	if !ok {
		return false, errors.NewError(errors.TokenWrong, nil)
	}
	val, err := CommonServiceApp.RedisGet(username.(string))
	if err != nil {
		log.Printf("redis get error:%v", err)
		return false, err
	} else if val == nil {
		return false, nil
	}
	return true, nil
}

func (u *UserService) SearchAttentionUser(username string, limit int, offset int, s *xorm.Session) (*[]string, error) {
	s = CommonServiceApp.SetSession(s)
	attention := make([]*model.Attention, 0)
	err := s.Cols("a_username").Where("username = ?", username).OrderBy("create_at").Limit(limit, offset).Find(&attention)
	if err != nil {
		log.Printf("select attention error:%v", err)
		return nil, err
	}
	users := make([]string, len(attention))
	for i := 0; i < len(attention); i++ {
		users[i] = attention[i].AUsername
	}
	return &users, nil
}

func (u *UserService) SearchUserByUsername(usernames *[]string, s *xorm.Session) (*[]*model.User, error) {
	s = CommonServiceApp.SetSession(s)
	users := make([]*model.User, 0)
	err := s.Cols("username", "nickname", "profile").In("username", *usernames).Find(&users)
	if err != nil {
		log.Printf("select user error: %v", err)
		return nil, err
	}
	return &users, nil
}
