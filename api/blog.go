package api

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/models"
	"backend-learning/hw2/pkg/request"
	"backend-learning/hw2/pkg/response"
	"backend-learning/hw2/service"
	"github.com/gin-gonic/gin"
)

var Blog = new(blogApi)

type blogApi struct {
}

func (b *blogApi) CreateTag(ctx *gin.Context) {
	// TODO 创建标签

	var params request.CreateTagRequest
	_ = ctx.ShouldBind(&params)

	if err := global.Validate.Struct(params); err != nil {
		response.Failed(err)
	}

	tag := &models.Tag{
		Name: params.Name,
	}

	_, err := service.Tag.CreateTag(tag)
	if err != nil {
		response.Failed(err)
		return
	}
	response.Success()
}

func (b blogApi) GetTagPostList(ctx *gin.Context) {
	// TODO 获取标签下所有的文章
}
