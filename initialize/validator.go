package initialize

import (
	"backend-learning/hw2/global"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	validate := validator.New()

	global.Validate = validate
	global.LOG.Info("初始化validator校验器完成")
}
