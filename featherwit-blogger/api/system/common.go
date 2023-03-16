package system

import (
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

	// todo 初始化client放到global中，考虑使用配置文件
	Endpoint := "https://oss-cn-shenzhen.aliyuncs.com"
	AccessKeyId := "myAccessKeyId"
	AccessKeySecret := "myAccessKeySecret"
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	bucket, err := client.Bucket("featherwit-blog-img")
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}

	// todo 识别文件后缀，生成一个随机、唯一的文件名，添加上原后缀
	err = bucket.PutObject("img/test2.png", img)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	imgUrl, err := bucket.SignURL("img/test2.png", oss.HTTPGet, 3600)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, imgUrl, c)
}
