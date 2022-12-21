package system

import (
	"featherwit-blogger/api"
	middle_ware "featherwit-blogger/middle-ware"
	"github.com/gin-gonic/gin"
)

type BlogRouter struct{}

func (b *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	blogRouter := Router.Group("blog")
	blogApi := api.ApiGroupApp.SystemApiGroup.BlogApi
	{
		blogRouter.POST("add", middle_ware.ConsumerToken(), blogApi.AddBlog)
		blogRouter.GET("/:limit/:offset", blogApi.SearchBlog)
		blogRouter.GET("id/:id", blogApi.GetBlogById)
	}
}
