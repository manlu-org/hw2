package main

import (
	"backend-learning/hw2/core"
	"backend-learning/hw2/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	core.InitViper()
	core.InitLog()
	initialize.Mysql()
	initialize.DBTables()
	initialize.InitValidator()

	core.Run()
}
