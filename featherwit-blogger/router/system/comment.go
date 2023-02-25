package system

import (
	"featherwit-blogger/api"
	middle_ware "featherwit-blogger/middle-ware"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (cr *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment")
	commentApi := api.ApiGroupApp.SystemApiGroup.CommentApi
	{
		commentRouter.POST("add", middle_ware.LoginToken(), commentApi.AddComment)
		commentRouter.DELETE(":commentId", middle_ware.LoginToken(), commentApi.DeleteComment)
		commentRouter.GET(":blogId/:limit/:offset", commentApi.GetComment)
	}
}
