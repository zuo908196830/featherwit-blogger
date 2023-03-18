package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/response"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

	Endpoint := global.GlobalConfig.AccessKey.Endpoint
	AccessKeyId := global.GlobalConfig.AccessKey.AccessKeyId
	AccessKeySecret := global.GlobalConfig.AccessKey.AccessKeySecret
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	bucket, err := client.Bucket(global.GlobalConfig.AccessKey.ImgBucket)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}

	// todo 识别文件后缀，生成一个随机、唯一的文件名，添加上原后缀
	err = bucket.PutObject("img/test6.png", img)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	imgUrl, err := bucket.SignURL("img/test6.png", oss.HTTPGet, 3600)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, imgUrl, c)
}
