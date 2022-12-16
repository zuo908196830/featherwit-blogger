package response

import (
	"dianping/model/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkResponse struct {
	RequestId string `json:"-"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message"`
	Detail interface{} `json:"detail"`
}

type ErrorResponse struct {
	RequestId string `json:"-"`
	Code int `json:"code"`
	Error Error `json:"error"`
}

func BuildOkResponse(code int, data interface{}, c *gin.Context)  {
	c.JSON(code, OkResponse{
		RequestId: c.Request.Header.Get("Request-Id"),
		Code:      0,
		Data:      data,
	})
}

func BuildErrorResult(code int, msg string, detail interface{}, c *gin.Context)  {
	err := Error{
		Message: msg,
		Detail:  detail,
	}
	httpCode := code
	if code < 100 || code > 999 {
		httpCode = http.StatusOK
	}
	c.JSON(httpCode, ErrorResponse{
		RequestId: c.Request.Header.Get("Request-Id"),
		Code:      code,
		Error:     err,
	})
}

func BuildErrorResponse(e error, c *gin.Context)  {
	if e == nil {
		BuildErrorResult(http.StatusInternalServerError, "insternal server error", "", c)
	} else if err, ok := e.(*errors.Error); ok {
		BuildErrorResult(int(err.Code), err.Msg, err.Detail, c)
	} else {
		BuildErrorResult(http.StatusInternalServerError, "insternal server error", "", c)
	}
}
