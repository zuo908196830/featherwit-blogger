package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"featherwit-blogger/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Printf("bind request param error: %v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	user, err := UserService.GetUserByUsername(login.Username, nil)
	if err != nil {
		log.Printf("get user error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if user == nil {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "username not exist"), c)
		return
	}
	if user.Password != login.Password {
		response.BuildErrorResponse(errors.NewError(errors.PasswordWrong, nil), c)
		return
	}
	err = CommonService.RedisSet(user.Username, user.Role)
	if err != nil {
		log.Printf("set session error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	err = CommonService.RedisSetTime(user.Username, 1800)
	if err != nil {
		log.Printf("set session error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	mp := make(map[string]interface{})
	mp["username"] = user.Username
	mp["nickname"] = user.Nickname
	token, err := utils.NewToken(mp)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, &response.Login{
		Username: user.Username,
		Nickname: user.Nickname,
		Token:    token,
	}, c)
}

func (u *UserApi) Logout(c *gin.Context) {
	username := CommonService.GetUsername(c)
	CommonService.RedisDelete(username)
	response.BuildOkResponse(0, nil, c)
}

func (u *UserApi) Register(c *gin.Context) {
	var register request.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		log.Printf("bind request param error: %v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	user, err := UserService.GetUserByUsername(register.Username, nil)
	if err != nil {
		log.Printf("select error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if user != nil {
		response.BuildErrorResponse(errors.NewError(errors.ResourceAlreadyExist, "username already exist"), c)
		return
	}
	err = UserService.AddUser(&model.User{
		Username:  register.Username,
		Password:  register.Password,
		Role:      register.Role,
		Nickname:  register.Nickname,
		Telephone: register.Telephone,
		Mail:      register.Mail,
		Profile:   register.Profile,
	}, nil)
	if err != nil {
		log.Printf("add user error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	err = CommonService.RedisSet(register.Username, register.Role)
	if err != nil {
		log.Printf("set session error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	err = CommonService.RedisSetTime(register.Username, 1800)
	if err != nil {
		log.Printf("set session error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	mp := make(map[string]interface{})
	mp["username"] = register.Username
	mp["nickname"] = register.Nickname
	token, err := utils.NewToken(mp)
	if err != nil {
		log.Printf("create token error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, response.Login{
		Username: register.Username,
		Nickname: register.Nickname,
		Token:    token,
	}, c)
}

func (u *UserApi) GetUser(c *gin.Context) {
	username := CommonService.GetUsername(c)
	user, err := UserService.GetUserByUsername(username, nil)
	if err != nil {
		log.Printf("get user error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if user == nil {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "user not exist"), c)
		return
	}
	resp := &response.UserData{
		Username:  user.Username,
		Role:      user.Role,
		Nickname:  user.Nickname,
		Telephone: user.Telephone,
		Mail:      user.Mail,
		Profile:   user.Profile,
	}
	response.BuildOkResponse(0, resp, c)
}

func (u *UserApi) LoginStatus(c *gin.Context) {
	ok, exists := c.Get("login-status")
	if exists {
		response.BuildOkResponse(0, ok.(bool), c)
	} else {
		response.BuildOkResponse(0, false, c)
	}
}

func (u *UserApi) TokenLogin(c *gin.Context) {
	value, _ := c.Get("User-Info")
	tkmp := value.(map[string]interface{})
	response.BuildOkResponse(0, tkmp, c)
}

func (u *UserApi) AttentionUser(c *gin.Context) {
	var param request.Page
	c.ShouldBindUri(&param)
	username := CommonService.GetUsername(c)
	usernames, err := UserService.SearchAttentionUser(username, param.Limit, param.Offset, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	users, err := UserService.SearchUserByUsername(usernames, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	attentionUsers := make([]response.AttentionUserData, len(*users))
	for i := 0; i < len(*users); i++ {
		attentionUsers[i].Username = (*users)[i].Username
		attentionUsers[i].Nickname = (*users)[i].Nickname
		attentionUsers[i].Profile = (*users)[i].Profile
	}
	response.BuildOkResponse(0, &attentionUsers, c)
}

func (u *UserApi) AddAttentionUser(c *gin.Context) {
	var param request.AddAttentionRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	session := global.DbEngine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		log.Printf("session begin error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	username := CommonService.GetUsername(c)
	if err := UserService.FansUpdate(param.AUsername, 1, session); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if err := UserService.AddAttentionUser(username, param.AUsername, session); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if err := session.Commit(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
}
