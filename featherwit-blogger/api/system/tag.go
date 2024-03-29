package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"strconv"

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
	for _, tagName := range param.TagNames {
		tags = append(tags, &model.Tag{
			Name: tagName,
		})
	}
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	n, err := TagService.AddTag(&tags, session)
	if err != nil {
		response.BuildErrorResponse(err, nil)
		return
	} else if n < int64(len(tags)) {
		response.BuildErrorResponse(errors.NewError(errors.ActionFail, nil), nil)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, &response.SearchTagResponse{Tags: tags}, c)
}

func (t *TagApi) AddTagBlog(c *gin.Context) {
	var param request.AddTagBlogRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	has, err := BlogService.BlogExist(param.BlogId, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	} else if !has {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
		return
	}
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if err := TagService.AddTagBlog(param.BlogId, &param.TagIds, session); err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (t *TagApi) DeleteTagBlog(c *gin.Context) {
	s := c.Param("blogId")
	username := CommonService.GetUsername(c)
	blogId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	blog, err := BlogService.GetBlogById(blogId, session)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if blog.Username != username {
		response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
		return
	}
	ok, err := TagService.DeleteTagBlogByBlogId(blogId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ActionFail, nil), c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (t *TagApi) SearchTag(c *gin.Context) {
	var page request.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.BuildErrorResponse(err, nil)
		return
	}
	res, err := TagService.SearchTag(page.Limit, page.Offset, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, *res, c)
}
