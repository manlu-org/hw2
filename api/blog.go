package api

import "github.com/gin-gonic/gin"

var Blog = new(blogApi)

type blogApi struct {
}

func (b *blogApi) CreateTag(ctx *gin.Context) {
	// TODO 创建标签
}

func (b blogApi) GetTagPostList(ctx *gin.Context) {
	// TODO 获取标签下所有的文章
}
