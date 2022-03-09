package initialize

import (
	"backend-learning/hw2/global"
	"backend-learning/hw2/models"
)

// 注册数据库表专用
func DBTables() {
	db := global.DB
	err := db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").
		AutoMigrate(
			&models.Tag{},
		)
	if err != nil {
		global.LOG.Debug(err)
		return
	}
	global.LOG.Debug("register table success")
}
