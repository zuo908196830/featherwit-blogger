package system

import "the-fool-blogger/service"

type ApiGroup struct {
	UserApi
}

var (
	UserService = new(service.UserService)
)
