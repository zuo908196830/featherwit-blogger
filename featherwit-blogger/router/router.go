package router

import (
	"featherwit-blogger/router/system"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	systemRouter system.SystemRouter
}

var RouterGroupApp = new(Routers)

func Init() *gin.Engine {
	Router := gin.New()
	outRouter := Router.Group("api")
	{
		RouterGroupApp.systemRouter.InitUserRouter(outRouter)
		RouterGroupApp.systemRouter.InitBlogRouter(outRouter)
	}
	return Router
}
