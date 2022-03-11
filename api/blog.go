package api

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/models"
	"backend-learning/hw2/pkg/request"
	"backend-learning/hw2/pkg/response"
	"backend-learning/hw2/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var Blog = new(blogApi)

type blogApi struct {
}

func (b *blogApi) CreateTag(ctx *gin.Context) {
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

func (b blogApi) GetAllTags(ctx *gin.Context) {

	tags, err := service.Tag.GetAllTags()
	if err != nil {
		response.FailWithMessage("系统错误")
		return
	}

	response.SuccessWithData(tags)
}

func (b blogApi) GetTag(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage("输入的ID有误")
		return
	}
	tag, err := service.Tag.GetTag(id)

	if err != nil {
		response.FailWithMessage("系统错误")
		return
	}

	response.SuccessWithData(tag)
}

func (b blogApi) CreatePost(ctx *gin.Context) {
	var req request.CreatePostRequest
	_ = ctx.ShouldBind(&req)

	post := &models.Post{
		Title:   req.Title,
		Content: req.Content,
	}

	_, err := service.Post.CreatePost(post)
	if err != nil {
		response.FailWithMessage("created error")
		return
	}
	response.Success()
}

func (b blogApi) GetPost(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage("输入的ID有误")
		return
	}

	post, err := service.Post.GetPost(id)
	if err != nil {
		response.FailWithMessage("未找到")
		return
	}

	response.SuccessWithData(post)
}

func (b blogApi) GetPostsByTag(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage("输入的ID有误")
		return
	}

	posts, err := service.Tag.GetPostByTagID(id)
	if err != nil {
		response.FailWithMessage(err.Error())
		return
	}
	response.SuccessWithData(posts)
}

func (b blogApi) GetAllPosts(ctx *gin.Context) {
	var req request.PageRequest
	_ = ctx.ShouldBind(&req)

	if req.PageSize == 0 {
		req.PageSize = 30
	}

	posts, total, err := service.Post.GetAllPosts(req.Page, req.PageSize)
	if err != nil {
		response.FailWithMessage("errors")
	}

	response.SuccessWithData(response.PostResp{
		List: posts,
		Meta: &response.Meta{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int(total),
		},
	})
}

func (b blogApi) ReplyPost(ctx *gin.Context) {
	var req request.CreateReplyRequest
	_ = ctx.ShouldBind(&req)

	if err := global.Validate.Struct(req); err != nil {
		response.Failed(err)
	}

	reply := &models.Reply{
		Content: req.Content,
		PostID:  uint(req.PostID),
	}

	err := service.Post.ReplyPost(reply)
	if err != nil {
		response.Failed(err)
		return
	}
	response.Success()
}

func (b blogApi) DelPost(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailWithMessage("输入的ID有误")
		return
	}

	err = service.Post.DelPost(id)
	if err != nil {
		response.FailWithMessage("delete error")
		return
	}

	response.Success()
}
