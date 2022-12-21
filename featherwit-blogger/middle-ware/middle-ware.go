package middle_ware

import (
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/response"
	"featherwit-blogger/service"
	"featherwit-blogger/utils"
	"github.com/gin-gonic/gin"
)

func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("User-Info")
		if token == "" {
			c.Next()
		} else {
			pkmp, err := utils.ParseToken(token)
			if err != nil {
				response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
				c.Abort()
				return
			}
			c.Set("User-Info", pkmp)
			ok, err := service.UserServiceApp.LoginStatus(pkmp)
			if err != nil {
				response.BuildErrorResponse(err, c)
				c.Abort()
				return
			}
			c.Set("login-status", ok)
			c.Next()
			if ok {
				service.CommonServiceApp.RedisSetTime(pkmp["username"].(string), 1800)
			}
		}
	}
}

func ConsumerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, exists := c.Get("login-status")
		if exists && ok.(bool) {
			c.Next()
		} else {
			response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
			c.Abort()
		}
	}
}
