package router

import (
	"backend-learning/hw2/api"
	"github.com/gin-gonic/gin"
)

func InitBlogGroup(router *gin.RouterGroup) {
	// TODO 创建标签
	router.POST("/tag", api.Blog.CreateTag)
	// TODO 获取标签列表

	// TODO 获取:id下的所有文章
	router.GET("/tag/:id/posts")
	// TODO 获取所有文章

	// TODO 获取id为:id的文章详情

	// TODO 创建一篇文章

}
