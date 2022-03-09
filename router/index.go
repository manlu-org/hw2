package router

import (
	"backend-learning/hw2/api"
	"github.com/gin-gonic/gin"
)

func InitBaseGroup(router *gin.RouterGroup) {
	router.GET("/hello", api.Hello.SayHello)
}
