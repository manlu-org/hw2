package api

import "github.com/gin-gonic/gin"

var Hello = new(helloApi)

type helloApi struct {
}

func (h *helloApi) SayHello(ctx *gin.Context) {
	ctx.JSON(200, "Hello World")
}
