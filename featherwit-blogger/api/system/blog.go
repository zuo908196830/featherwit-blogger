package system

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"github.com/gin-gonic/gin"
	"log"
)

type BlogApi struct{}

func (b *BlogApi) AddBlog(c *gin.Context) {
	var param request.BlogRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("bad request error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	get, _ := c.Get("User-Info")
	tkmp := get.(map[string]interface{})
	username, ok := tkmp["username"].(string)
	user, err := UserService.GetUserByUsername(username)
	if err != nil {
		log.Printf("select user error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	if user == nil {
		log.Printf("username is not exist")
		response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
		return
	}
	if !ok {
		response.BuildErrorResponse(errors.NewError(errors.TokenWrong, "username is not exist"), c)
		return
	}
	blog := &model.Blob{
		Username:    username,
		Title:       param.Title,
		Content:     param.Content,
		Views:       0,
		CommonCount: 0,
		LikeCount:   0,
	}
	err = BlogService.AddBlog(blog)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, response.AddBlogResponse{
		ID:       blog.ID,
		CreateAt: blog.CreateAt,
		UpdateAt: blog.UpdateAt,
	}, c)
}

func (b *BlogApi) SearchBlog(c *gin.Context) {
	var param request.SearchBlogRequest
	if err := c.ShouldBindUri(&param); err != nil {
		log.Printf("bind uri error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	blogs, err := BlogService.SearchBlog(param.Limit, param.Offset)
	if err != nil {
		log.Printf("search blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, blogs, c)
}

func (b *BlogApi) GetBlogById(c *gin.Context) {
	var bid request.BlogIdRequest
	if err := c.ShouldBindUri(&bid); err != nil {
		log.Printf("bind uri error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	blog, err := BlogService.GetBlogById(bid.ID)
	if err != nil {
		log.Printf("select by id error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if blog == nil {
		log.Printf("blog is not exist")
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	response.BuildOkResponse(0, blog, c)
}

func (b *BlogApi) UpdateBlog(c *gin.Context) {
	var param request.UpdateBlogRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("bind update blog param error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	val, _ := c.Get("User-Info")
	tkmp := val.(map[string]interface{})
	username := tkmp["username"].(string)
	exist, err := BlogService.BlogExist(param.ID, username)
	if err != nil {
		log.Printf("get blog error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if !exist {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	err = BlogService.UpdateBlog(&model.Blob{
		ID:       param.ID,
		Username: username,
		Title:    param.Title,
		Content:  param.Content,
	})
	if err != nil {
		log.Printf("update blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, nil, c)
}
