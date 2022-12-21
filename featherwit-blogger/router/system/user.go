package system

import (
	"featherwit-blogger/api"
	middle_ware "featherwit-blogger/middle-ware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	UserApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", UserApi.Login)
		userRouter.POST("register", UserApi.Register)
		userRouter.GET("logout", middle_ware.ConsumerToken(), UserApi.Logout)
	}
}
