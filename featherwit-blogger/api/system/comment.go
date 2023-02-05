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
	ok, err := BlogService.BlogExist(param.BlogId, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	if param.ParentId >= 0 {
		ok, err = CommentService.CommentExist(param.ParentId, nil)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		} else if !ok {
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "parent comment is not exist"), c)
			return
		}
	}
	if param.ReplyId >= 0 {
		ok, err = CommentService.CommentExist(param.ReplyId, nil)
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
	err = CommentService.AddComment(comment, session)
	if err != nil {
		log.Printf("database error:%v", err)
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	//把blog和一级评论中的评论数加1
	ok, err = BlogService.UpdateCommentCount(param.BlogId, 1, session)
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
		ok, err := CommentService.UpdateCommentCount(param.ParentId, 1, session)
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
	s := c.Param("commentId")
	commentId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("strconv error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	mp, _ := c.Get("User-Info")
	tkmp := mp.(map[string]interface{})
	username := tkmp["username"].(string)
	comment, err := CommentService.GetCommentById(commentId, nil)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	} else if comment == nil {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
		return
	}
	if comment.Username != username {
		response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
		return
	}
	session := global.DbEngine.NewSession()
	session.Begin()
	defer session.Close()
	// 如果是二级评论则给一级评论和Blog评论数减1
	if comment.ParentId >= 0 {
		b, err := CommentService.UpdateCommentCount(comment.ParentId, -1, session)
		if err != nil {
			session.Rollback()
			response.BuildErrorResponse(err, c)
			return
		} else if !b {
			session.Rollback()
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
			return
		}
	}
	// 先删除子评论，再删除自身
	children, err := CommentService.SearchCommentByParentId(commentId, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	deleteList := make([]int64, len(children))
	for i := 0; i < len(children); i++ {
		deleteList[i] = children[i].ID
	}
	deleteList = append(deleteList, commentId)
	err = CommentService.DeleteCommentById(deleteList, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	b, err := BlogService.UpdateCommentCount(comment.BlogId, -len(deleteList), session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	} else if !b {
		session.Rollback()
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, nil), c)
		return
	}
	session.Commit()
	response.BuildOkResponse(0, nil, c)
}

func (ca *CommentApi) GetComment(c *gin.Context) {
	// s := c.Param("blogId")
	// blogId, err := strconv.ParseInt(s, 10, 64)
	// if err != nil {
	// 	log.Printf("strconv error:%v", err)
	// 	response.BuildErrorResponse(err, c)
	// 	return
	// }
	// comments, err := CommentService.GetCommentByBlogId(blogId)
	// if err != nil {
	// 	response.BuildErrorResponse(err, c)
	// 	return
	// }

}
