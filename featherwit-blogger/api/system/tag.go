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
	TagService.MutilAddTag(param.Tags, "", "", 1, &tags)
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
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
	if err := TagService.AddTagBlog(param.BlogId, &param.Tags, session); err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (t *TagApi) DeleteTagBlog(c *gin.Context) {
	s := c.Param("blogId")
	val, _ := c.Get("User-Info")
	tkmp := val.(map[string]interface{})
	username := tkmp["username"].(string)
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
	res, err := TagService.SearchTag(nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, *res, c)
}
