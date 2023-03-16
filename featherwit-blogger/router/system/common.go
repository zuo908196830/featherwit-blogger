package system

import (
	"featherwit-blogger/api"

	"github.com/gin-gonic/gin"
)

type CommonRouter struct{}

func (c *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
	commonRouter := Router.Group("common")
	commonApi := api.ApiGroupApp.SystemApiGroup.CommonApi
	{
		commonRouter.POST("upload", commonApi.Upload)
	}
}
