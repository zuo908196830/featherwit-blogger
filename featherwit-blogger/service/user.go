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
	s = CommentServiceApp.SetSession(s)
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
	s = CommentServiceApp.SetSession(s)
	_, err := s.Insert(user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserService) LoginStatus(user map[string]interface{}, s *xorm.Session) (bool, error) {
	s = CommentServiceApp.SetSession(s)
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
