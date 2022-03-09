package middleware

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

var Middleware = new(middleware)

type middleware struct {
}

// 全局异常处理中间件
func (m *middleware) Exception(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 将异常写入日志
			//global.LOG.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			if resp, ok := err.(response.Resp); ok {
				response.JSON(c, http.StatusOK, resp)
				c.Abort()
				return
			}
			// 服务器异常
			resp := response.Resp{
				Code: response.InternalServerError,
				Data: map[string]interface{}{},
				Msg:  response.CustomError[response.InternalServerError],
			}
			global.LOG.Error(fmt.Sprintf("[Exception]未知异常: %v\n堆栈信息: %v", err, string(debug.Stack())))
			// 以json方式写入响应
			response.JSON(c, http.StatusOK, resp)
			c.Abort()
			return
		}
	}()
	c.Next()
}
