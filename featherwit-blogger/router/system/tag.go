package system

import (
	"featherwit-blogger/api"

	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (t *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("tag")
	tagApi := api.ApiGroupApp.SystemApiGroup.TagApi
	{
		tagRouter.POST("add/tree", tagApi.AddTag)
	}
}
