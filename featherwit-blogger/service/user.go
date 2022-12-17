package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
)

type UserService struct{}

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
