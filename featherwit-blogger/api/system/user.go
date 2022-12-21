package system

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"featherwit-blogger/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Printf("bind request param error: %v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	user, err := UserService.GetUserByUsername(login.Username)
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
	mp["role"] = user.Role
	mp["nickname"] = user.Nickname
	token, err := utils.NewToken(mp)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, response.Login{
		Username: user.Username,
		Nickname: user.Nickname,
		Token:    token,
	}, c)
}

func (u *UserApi) Logout(c *gin.Context) {
	get, _ := c.Get("User-Info")
	tkmp := get.(map[string]interface{})
	username := tkmp["username"]
	CommonService.RedisDelete(username.(string))
	response.BuildOkResponse(0, nil, c)
}

func (u *UserApi) Register(c *gin.Context) {
	var register request.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		log.Printf("bind request param error: %v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	user, err := UserService.GetUserByUsername(register.Username)
	if err != nil {
		log.Printf("select error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if user != nil {
		response.BuildErrorResponse(errors.NewError(errors.ResourceAlreadyExist, "username already exist"), c)
		return
	}
	err = UserService.AddUser(&model.User{
		Username: register.Username,
		Password: register.Password,
		Role:     register.Role,
		Nickname: register.Nickname,
	})
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
	mp["role"] = register.Role
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
