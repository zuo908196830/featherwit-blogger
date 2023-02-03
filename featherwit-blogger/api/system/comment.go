package system

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/request"
	"featherwit-blogger/model/response"
	"log"

	"github.com/gin-gonic/gin"
)

type CommentApi struct{}

// todo 事务RollBack和close出错时的处理
func (ca *CommentApi) AddComment(c *gin.Context) {
	var param request.AddComment
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Printf("bad request err:%v", err)
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	mp, _ := c.Get("User-Info")
	tkmp := mp.(map[string]interface{})
	username := tkmp["username"]
	comment := &model.Comment{
		Username: username.(string),
		BlogId:   param.BlogId,
		ParentId: param.ParentId,
		ReplyId:  param.ReplyId,
		Content:  param.Content,
	}
	// 添加BlogId ParaentId ReplyId存在性校验
	ok, err := BlogService.BlogExist(param.BlogId)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	if param.ParentId >= 0 {
		ok, err = CommentService.CommentExist(param.ParentId)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		} else if !ok {
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "parent comment is not exist"), c)
			return
		}
	}
	if param.ReplyId >= 0 {
		ok, err = CommentService.CommentExist(param.ReplyId)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		} else if !ok {
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "reply comment is not exist"), c)
			return
		}
	}
	// 开始添加评论
	session := global.DbEngine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		log.Printf("open transaction faild")
		response.BuildErrorResponse(err, c)
		return
	}
	err = CommentService.AddComment(comment)
	if err != nil {
		log.Printf("database error:%v", err)
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	//把blog和一级评论中的评论数加1
	ok, err = BlogService.ContentCountPlus1(param.BlogId)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
		return
	}
	if param.ParentId >= 0 {
		ok, err := CommentService.ContentCountPlus1(param.ParentId)
		if err != nil {
			session.Rollback()
			response.BuildErrorResponse(err, c)
			return
		} else if !ok {
			session.Rollback()
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
			return
		}
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (ca *CommentApi) DeleteComment(c *gin.Context) {

}
