package system

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/model/response"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	UserService.Login()
	response.BuildOkResponse(0, "login", c)
}
