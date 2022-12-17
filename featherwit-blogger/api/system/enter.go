package system

import "featherwit-blogger/service"

type ApiGroup struct {
	UserApi
}

var (
	UserService = new(service.UserService)
)
