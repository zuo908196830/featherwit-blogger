package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	if !ok {
		response.BuildErrorResponse(errors.NewError(errors.TokenWrong, "username is not exist"), c)
		return
	}
	user, err := UserService.GetUserByUsername(username, nil)
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
	blog := &model.Blog{
		Username: user.Username,
		Title:    param.Title,
		Content:  param.Content,
		Profile:  param.Profile,
	}
	err = BlogService.AddBlog(blog, nil)
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
	var param request.Page
	if err := c.ShouldBindUri(&param); err != nil {
		log.Printf("bind uri error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	blogs, err := BlogService.SearchBlog(param.Limit, param.Offset, nil)
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
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	blog, err := BlogService.GetBlogById(bid.ID, nil)
	if err != nil {
		log.Printf("select by id error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if blog == nil {
		log.Printf("blog is not exist")
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	if err := session.Commit(); err != nil {
		response.BuildErrorResponse(err, c)
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
	username := CommonService.GetUsername(c)
	exist, err := BlogService.BlogExist(param.ID, nil)
	if err != nil {
		log.Printf("get blog error: %v", err)
		response.BuildErrorResponse(err, c)
		return
	} else if !exist {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	err = BlogService.UpdateBlog(&model.Blog{
		ID:       param.ID,
		Username: username,
		Title:    param.Title,
		Content:  param.Content,
	}, nil)
	if err != nil {
		log.Printf("update blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, nil, c)
}

func (b *BlogApi) GetBlogCount(c *gin.Context) {
	n, err := BlogService.BlogCount(nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, &response.BlogCountResponse{
		Total: n,
	}, c)
}

func (b *BlogApi) DeleteBlog(c *gin.Context) {
	s := c.Param("blogId")
	blogId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	username := CommonService.GetUsername(c)
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	// 验证权限
	blog, err := BlogService.GetBlogById(blogId, session)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if blog.Username != username {
		response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
		return
	}
	// 删除评论
	ok, err := CommentService.DeleteCommentByBlogId(blogId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ActionFail, nil), c)
		return
	}
	// 删除blog与tag之间的绑定
	ok, err = TagService.DeleteTagBlogByBlogId(blogId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ActionFail, nil), c)
		return
	}
	// 删除blog
	err = BlogService.DeleteBlogById(blogId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (b *BlogApi) UpdateCover(c *gin.Context) {
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
	s := c.Param("blogId")
	blogId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	username := CommonService.GetUsername(c)
	blog, err := BlogService.GetBlogById(blogId, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	if blog.Username != username {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
		return
	}
	blog.Cover = url
	err = BlogService.UpdateBlog(blog, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, nil, c)
}
