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
		userRouter.GET("logout", middle_ware.LoginToken(), UserApi.Logout)
		userRouter.GET("data", middle_ware.LoginToken(), UserApi.GetUser)
		userRouter.GET("status", UserApi.LoginStatus)
		userRouter.GET("token/login", middle_ware.LoginToken(), UserApi.TokenLogin)
		userRouter.GET("attention/:limit/:offset", middle_ware.LoginToken(), UserApi.AttentionUser)
		userRouter.POST("attention/add", middle_ware.LoginToken(), UserApi.AddAttentionUser)
	}
}
