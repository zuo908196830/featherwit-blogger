package middle_ware

import (
	"featherwit-blogger/model/errors"
	"featherwit-blogger/model/response"
	"featherwit-blogger/service"
	"featherwit-blogger/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
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
			ok, err := service.UserServiceApp.LoginStatus(pkmp, nil)
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

func LoginToken() gin.HandlerFunc {
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

func AdministratorsToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		get, _ := c.Get("User-Info")
		tkmp := get.(map[string]interface{})
		role, err := service.CommonServiceApp.RedisGet(tkmp["username"].(string))
		if err != nil {
			response.BuildErrorResponse(err, c)
			c.Abort()
			return
		} else if role != 1 {
			response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func AddBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("User-Info")
		if exists {
			c.Next()
		} else {
			response.BuildErrorResponse(errors.NewError(errors.Unauthorized, nil), c)
			c.Abort()
			return
		}
	}
}
