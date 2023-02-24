package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (ta *TagApi) AddTag(c *gin.Context) {
	var param request.AddTagRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	tags := make([]*model.Tag, 0)
	TagService.MutilAddTag(param.Tags, "", "", 1, &tags)
	session := global.DbEngine.NewSession()
	defer session.Close()
	session.Begin()
	n, err := TagService.AddTag(&tags, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	if int(n) != len(tags) {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ResourceAlreadyExist, nil), c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}
