package service

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"log"

	"xorm.io/xorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (u *UserService) UserExist(username string, s *xorm.Session) (bool, error) {
	s = CommonServiceApp.SetSession(s)
	has, err := s.Where("username = ?", username).Exist(new(model.User))
	if err != nil {
		return false, err
	}
	return has, nil
}

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

func (u *UserService) AddAttentionUser(username string, aUsername string, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	att := &model.Attention{
		Username:  username,
		AUsername: aUsername,
	}
	if exist, err := s.Exist(att); err != nil {
		return err
	} else if exist {
		return errors.NewError(errors.ResourceAlreadyExist, nil)
	}
	if _, err := s.Insert(att); err != nil {
		return err
	}
	return nil
}

func (u *UserService) FansUpdate(username string, num int, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	var user model.User
	if has, err := s.Cols("fans_count").Where("username = ?", username).ForUpdate().Get(&user); err != nil {
		return err
	} else if !has {
		return errors.NewError(errors.ResourceNotExist, "aUsername not exist")
	}
	if _, err := s.Cols("fans_count").Where("username = ?", username).Update(&model.User{
		FansCount: user.FansCount + int64(num),
	}); err != nil {
		return err
	}
	return nil
}

func (u *UserService) ConcernUser(username string, concernUsername string, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	concern := &model.ConcernList{
		Username:        username,
		ConcernUsername: concernUsername,
	}
	exist, err := s.Exist(concern)
	if err != nil {
		log.Printf("select concern_list error:%v", err)
		return err
	} else if exist {
		return nil
	}
	_, err = s.Insert(concern)
	if err != nil {
		log.Printf("Insert concern_list error:%v", err)
		return err
	}
	return nil
}

func (u *UserService) UnConcernUser(username string, concernUsername string, s *xorm.Session) error {
	s = CommonServiceApp.SetSession(s)
	concern := &model.ConcernList{
		Username:        username,
		ConcernUsername: concernUsername,
	}
	exist, err := s.Exist(concern)
	if err != nil {
		log.Printf("select concern_list error:%v", err)
		return err
	} else if !exist {
		return errors.NewError(errors.ResourceNotExist, nil)
	}
	_, err = s.Delete(concern)
	if err != nil {
		log.Printf("delete concern_list error:%v", err)
		return err
	}
	return nil
}

func (u *UserService) SearchConcernUser(username string, limit int, offset int, s *xorm.Session) ([]*model.User, error) {
	s = CommonServiceApp.SetSession(s)
	concerns := make([]*model.ConcernList, 0)
	err := s.Where("username = ?", username).Limit(limit, offset).Decr("create_at").Find(&concerns)
	if err != nil {
		log.Printf("select concern_list error:%v", err)
		return nil, err
	}
	usernames := make([]string, len(concerns))
	for i := 0; i < len(concerns); i++ {
		usernames[i] = concerns[i].ConcernUsername
	}
	users := make([]*model.User, 0)
	err = s.In("username", usernames).Find(&users)
	if err != nil {
		log.Printf("select concern_list error:%v", err)
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetNicknameAndCover(username string, s *xorm.Session) (string, string, error) {
	s = CommonServiceApp.SetSession(s)
	var user model.User
	if has, err := s.Cols("nickname", "headshot").Where("username = ?", username).Get(&user); err != nil {
		log.Printf("select user error:%v", err)
		return "", "", err
	} else if !has {
		return "", "", errors.NewError(errors.ResourceNotExist, nil)
	}
	if user.Nickname == "" {
		return username, user.Headshot, nil
	}
	return user.Nickname, user.Headshot, nil
}
