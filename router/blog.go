package router

import (
	"backend-learning/hw2/api"
	"backend-learning/hw2/middleware"
	"github.com/gin-gonic/gin"
)

func InitBlogGroup(router *gin.RouterGroup) {
	// 创建标签
	router.POST("/tag", api.Blog.CreateTag)
	// 获取标签列表
	router.GET("/tags", api.Blog.GetAllTags)
	// 获取标签
	router.GET("/tag/:id", api.Blog.GetTag)
	// 创建文章
	router.POST("/post", api.Blog.CreatePost)
	// 文章详情
	router.GET("/post/:id", api.Blog.GetPost)
	// 获取某标签下的文章
	router.GET("/tag/:id/posts", api.Blog.GetPostsByTag)
	// 获取所有文章
	router.GET("/posts", api.Blog.GetAllPosts)
	// 回复文章
	router.POST("/post/:id/reply", api.Blog.ReplyPost)

	router.POST("/admin/login", middleware.GinJWTMiddleware.LoginHandler)

	auth := router.Use(middleware.GinJWTMiddleware.MiddlewareFunc())
	// 删除文章
	auth.DELETE("/post/:id/delete", api.Blog.DelPost)
}
