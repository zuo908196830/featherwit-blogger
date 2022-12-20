package system

import "featherwit-blogger/service"

type ApiGroup struct {
	UserApi
	BlogApi
}

var (
	CommonService = new(service.CommonService)
	UserService   = new(service.UserService)
	BlogService   = new(service.BlogService)
)
