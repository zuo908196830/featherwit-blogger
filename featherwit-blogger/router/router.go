package router

import (
	middle_ware "featherwit-blogger/middle-ware"
	"featherwit-blogger/router/system"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	systemRouter system.SystemRouter
}

var RouterGroupApp = new(Routers)

func Init() *gin.Engine {
	Router := gin.New()
	Router.Use(middle_ware.Cors(), middle_ware.TokenMiddleWare())
	outRouter := Router.Group("api")
	{
		RouterGroupApp.systemRouter.InitUserRouter(outRouter)
		RouterGroupApp.systemRouter.InitBlogRouter(outRouter)
		RouterGroupApp.systemRouter.InitCommentRouter(outRouter)
		RouterGroupApp.systemRouter.InitTagRouter(outRouter)
	}
	return Router
}
