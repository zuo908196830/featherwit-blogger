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
		Cover:    global.GlobalConfig.DefaultBlogCover,
	}
	if param.Cover != "" {
		blog.Cover = param.Cover
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
	var param request.SearchBlogRequest
	if err := c.ShouldBindQuery(&param); err != nil {
		log.Printf("bind JSON error:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	if param.Order < 1 || param.Order > 4 {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	if param.TagId != 0 {
		tagId := []int64{param.TagId}
		go TagService.AddSearchCount(tagId, nil)
	}
	blogs, err := BlogService.SearchBlog(&param, param.Limit, param.Offset, nil)
	if err != nil {
		log.Printf("search blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	total, err := BlogService.BlogCount(&param, nil)
	if err != nil {
		log.Printf("search blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, &response.SearchBlogs{
		Total: total,
		Blogs: blogs,
	}, c)
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
	ret := new(response.GetBlog)
	ret.Blog = blog
	if nickname, headshot, err := UserService.GetNicknameAndCover(blog.Username, session); err != nil {
		response.BuildErrorResponse(err, c)
	} else {
		ret.User = &response.UserShow{
			Nickname: nickname,
			Headshot: headshot,
		}
	}
	go b.TagAdd1(bid.ID)
	if err := session.Commit(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, ret, c)
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
		Profile:  param.Profile,
	}, nil)
	if err != nil {
		log.Printf("update blog error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, nil, c)
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

func (b *BlogApi) TagAdd1(blogId int64) {
	tagIds, err := TagService.SearchTagByBlogId(blogId, nil)
	if err != nil {
		return
	}
	TagService.AddSearchCount(tagIds, nil)
}

func (b *BlogApi) StarBlog(c *gin.Context) {
	username := CommonService.GetUsername(c)
	var param request.StarRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	go BlogService.BlogLikeCountAddN(param.BlogId, 1, nil)
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	for _, blogId := range param.BlogId {
		exist, err := BlogService.BlogExist(blogId, session)
		if err != nil {
			session.Rollback()
			response.BuildErrorResponse(err, c)
			return
		} else if !exist {
			session.Rollback()
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
			return
		}
	}
	n, err := BlogService.StarBlogs(username, param.BlogId, session)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, n, c)
}

func (b *BlogApi) UnStarBlog(c *gin.Context) {
	username := CommonService.GetUsername(c)
	var param request.StarRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	go BlogService.BlogLikeCountAddN(param.BlogId, -1, nil)
	session := global.DbEngine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	n, err := BlogService.UnStarBlogs(username, param.BlogId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, n, c)
}

func (b BlogApi) SearchStarBlog(c *gin.Context) {
	username := CommonService.GetUsername(c)
	blogs, err := BlogService.SearchBlogByStar(username, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(0, blogs, c)
}
