package system

import (
	"featherwit-blogger/api"
	middle_ware "featherwit-blogger/middle-ware"

	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (t *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagRouter := Router.Group("tag")
	tagApi := api.ApiGroupApp.SystemApiGroup.TagApi
	{
		tagRouter.POST("add/tree", middle_ware.LoginToken(), middle_ware.AdministratorsToken(), tagApi.AddTag)
		tagRouter.POST("add/tag/blog", middle_ware.LoginToken(), tagApi.AddTagBlog)
		tagRouter.GET("search", tagApi.SearchTag)
	}
}
