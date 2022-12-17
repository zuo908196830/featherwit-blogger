package system

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
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
	if user.Role == 0 {
		//todo 生成管理员token
		response.BuildOkResponse(0, response.Login{
			Username: user.Username,
			Nickname: user.Nickname,
			Token:    "0000",
		}, c)
	} else if user.Role == 1 {
		//todo 生成普通用户token
		response.BuildOkResponse(0, response.Login{
			Username: user.Username,
			Nickname: user.Nickname,
			Token:    "0000",
		}, c)
	}
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
	response.BuildOkResponse(0, nil, c)
}
