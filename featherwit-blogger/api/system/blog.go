package system

import (
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"featherwit-blogger/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type BlogApi struct{}

func (b *BlogApi) AddBlog(c *gin.Context) {
	token := c.GetHeader("User-Info")
	if token == "" {
		log.Printf("token is not exist")
		response.BuildErrorResponse(errors.NewError(errors.Unauthorized, "token is not exist"), c)
		return
	}
	var param request.BlogRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("bad request error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	tkmp, err := utils.ParseToken(token)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
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
