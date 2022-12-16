package system

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/api/system"
)

type ShopRouter struct {}

func (s *ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	ShopApi := system.ShopApi{}
	{
		Router.GET("/:id", ShopApi.GetById)
	}
}
