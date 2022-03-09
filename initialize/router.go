package initialize

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/middleware"
	"backend-learning/hw2/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.New()
	// 跨域
	Router.Use(middleware.Cors())
	Router.Use(middleware.Middleware.Exception)
	global.LOG.Debug("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	apiGroup := Router.Group("/api")
	router.InitBaseGroup(apiGroup)
	router.InitBlogGroup(apiGroup)
	return Router
}
