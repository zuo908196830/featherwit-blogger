package system

import "featherwit-blogger/service"

type ApiGroup struct {
	UserApi
}

var (
	CommonService = new(service.CommonService)
	UserService   = new(service.UserService)
)
