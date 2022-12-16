package system

import (
	"github.com/gin-gonic/gin"
	"the-fool-blogger/model/response"
)

type ShopTypeApi struct{}

func (s *ShopTypeApi) ShopTypeList(c *gin.Context) {
	response.BuildOkResponse(0, "ok", c)
}
