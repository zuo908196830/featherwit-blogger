package system

import (
	"featherwit-blogger/model/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	UserService.Login()
	response.BuildOkResponse(0, "login", c)
}
