package system

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/api"
)

type ShouTypeRouter struct {}

func (s *ShouTypeRouter) InitShopTypeRouter(Router *gin.RouterGroup) {
	ShopTypeApi := api.ApiGroupApp.SystemApiGroup
	Router.GET("/shop-type/list", ShopTypeApi.ShopTypeList)
}
