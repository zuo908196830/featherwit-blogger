package system

import (
	"featherwit-blogger/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	UserApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", UserApi.Login)
		userRouter.POST("register", UserApi.Register)
		userRouter.GET("logout", UserApi.Logout)
	}
}
