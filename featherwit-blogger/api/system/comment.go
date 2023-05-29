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
	username := CommonService.GetUsername(c)
	comment := &model.Comment{
		Username: username,
		BlogId:   param.BlogId,
		ParentId: param.ParentId,
		ReplyId:  param.ReplyId,
		Content:  param.Content,
	}
	// 添加BlogId ParaentId ReplyId存在性校验
	if ok, err := BlogService.BlogExist(param.BlogId, nil); err != nil {
		response.BuildErrorResponse(err, c)
		return
	} else if !ok {
		response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "blog is not exist"), c)
		return
	}
	if param.ParentId > 0 {
		ok, err := CommentService.CommentExist(param.ParentId, nil)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		} else if !ok {
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "parent comment is not exist"), c)
			return
		}
	}
	replyUsername := ""
	if param.ReplyId > 0 {
		replyComment, err := CommentService.GetCommentById(param.ReplyId, nil)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		} else if replyComment == nil {
			response.BuildErrorResponse(errors.NewError(errors.ResourceNotExist, "reply comment is not exist"), c)
			return
		}
		replyUsername = replyComment.Username
	}
	// 开始添加评论
	session := global.DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Printf("open transaction faild")
		response.BuildErrorResponse(err, c)
		return
	}
	cm, err := CommentService.AddComment(comment, session)
	if err != nil {
		log.Printf("database error:%v", err)
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	//把blog和一级评论中的评论数加1
	if ok, err := BlogService.UpdateCommentCount(param.BlogId, 1, session); err != nil {
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
	nickname, headshot, err := UserService.GetNicknameAndCover(username, session)
	if err != nil {
		session.Rollback()
		response.BuildErrorResponse(err, c)
		return
	}
	session.Commit()
	if param.ReplyId > 0 {
		if replyNickname, _, err := UserService.GetNicknameAndCover(replyUsername, nil); err != nil {
			session.Rollback()
			response.BuildErrorResponse(err, c)
			return
		} else {
			response.BuildOkResponse(0, &response.ChildrenComment{
				Comment: cm,
				User: &response.UserShow{
					Nickname: nickname,
					Headshot: headshot,
				},
				ReplyNickname: replyNickname,
			}, c)
			return
		}
	}
	response.BuildOkResponse(0, &response.CommentsTree{
		Comment:         cm,
		ChildrenCount:   0,
		ChildrenComment: make([]*response.ChildrenComment, 0),
		User: &response.UserShow{
			Nickname: nickname,
			Headshot: headshot,
		},
	}, c)
}

func (ca *CommentApi) DeleteComment(c *gin.Context) {
	s := c.Param("commentId")
	commentId, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("strconv error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	username := CommonService.GetUsername(c)
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
	comments := []int64{commentId}
	children, err := CommentService.SearchCommentByParentId(comments, session)
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
	var param request.GetBlogComment
	err := c.ShouldBindUri(&param)
	if err != nil {
		response.BuildErrorResponse(errors.NewError(errors.BadRequest, nil), c)
		return
	}
	if err != nil {
		log.Printf("strconv error:%v", err)
		response.BuildErrorResponse(err, c)
		return
	}
	session := global.DbEngine.NewSession()
	defer session.Close()
	comments, err := CommentService.GetCommentByBlogId(param.BlogId, param.Limit, param.Offset, session)
	if err != nil {
		response.BuildErrorResponse(err, c)
		return
	}
	// todo 通过一级评论去查找对应的二级评论
	commentTree := make([]*response.CommentsTree, 0)
	for _, comment := range comments {
		children, err := CommentService.GetCommentByParentID(comment.ID, session)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		}
		nickname, headshot, err := UserService.GetNicknameAndCover(comment.Username, nil)
		if err != nil {
			response.BuildErrorResponse(err, c)
			return
		}
		commentTree = append(commentTree, &response.CommentsTree{
			Comment:       comment,
			ChildrenCount: len(children),
			User: &response.UserShow{
				Nickname: nickname,
				Headshot: headshot,
			},
			ChildrenComment: children,
		})
	}
	response.BuildOkResponse(0, &response.Comments{
		Count:        len(commentTree),
		CommentsTree: commentTree,
	}, c)
}
