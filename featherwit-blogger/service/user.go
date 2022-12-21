package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"log"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	user := new(model.User)
	ok, err := global.DbEngine.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	} else {
		return user, nil
	}
}

func (u *UserService) AddUser(user *model.User) error {
	_, err := global.DbEngine.Insert(user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserService) LoginStatus(user map[string]interface{}) (bool, error) {
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
