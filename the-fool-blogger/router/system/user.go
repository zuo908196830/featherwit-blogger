package system

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/api"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	UserApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", UserApi.Login)
	}
}
