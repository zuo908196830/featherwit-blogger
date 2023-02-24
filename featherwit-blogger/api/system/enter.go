package system

import "featherwit-blogger/service"

type ApiGroup struct {
	UserApi
	BlogApi
	CommentApi
	TagApi
}

var (
	CommonService  = service.CommonServiceApp
	UserService    = service.UserServiceApp
	BlogService    = service.BlogServiceApp
	CommentService = service.CommentServiceApp
	TagService     = service.TagServiceApp
)
