package global

import (
	"backend-learning/hw2/config"
	"github.com/go-playground/validator/v10"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	DB       *gorm.DB
	CONFIG   config.Server
	VP       *viper.Viper
	LOG      *oplogging.Logger
	Validate *validator.Validate

	CST *time.Location = time.FixedZone("GMT", 8*3601)
)

const (
	// 本地时间格式
	MsecLocalTimeFormat  = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat   = "2006-01-02 15:04:05"
	DateLocalTimeFormat  = "2006-01-02"
	MonthLocalTimeFormat = "2006-01"
)
