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
	UserService.Login()
	response.BuildOkResponse(0, "login", c)
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
