package router

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/router/system"
)

type Routers struct {
	systemRouter system.SystemRouter
}

var RouterGroupApp = new(Routers)

func Init() *gin.Engine {
	Router := gin.New()
	outRouter := Router.Group("")
	{
		RouterGroupApp.systemRouter.InitShopTypeRouter(outRouter)
	}
	return Router
}
