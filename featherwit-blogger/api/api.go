package api

import "featherwit-blogger/api/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
