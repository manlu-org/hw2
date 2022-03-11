package main

import (
	"backend-learning/hw2/core"
	"backend-learning/hw2/initialize"
	"backend-learning/hw2/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	core.InitViper()
	core.InitLog()
	initialize.Mysql()
	initialize.DBTables()
	initialize.InitValidator()
	middleware.InitJWT()
	core.Run()
}
