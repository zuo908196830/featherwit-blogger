package system

import (
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/response"

	"github.com/gin-gonic/gin"
)

type CommonApi struct{}

func (ca *CommonApi) Upload(c *gin.Context) {
	img, headers, err := c.Request.FormFile("image")
	if err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	//headers.Size 获取文件大小
	if headers.Size > 1024*1024*16 {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, "file to big"), c)
		return
	}

	url, err := CommonService.UploadImg(&img, headers)
	if err != nil {
		response.BuildErrorResponse(err, c)
	}
	response.BuildOkResponse(0, url, c)
}
