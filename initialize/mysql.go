package initialize

import (
	"backend-learning/hw2/global"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Mysql() {
	admin := global.CONFIG.Mysql
	dsn := admin.Username + ":" + admin.Password + "@(" + admin.Path + ")/" + admin.Dbname + "?" + admin.Config

	global.LOG.Info("数据库连接DSN: ", dsn)
	init := false
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20)*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("初始化mysql异常: 连接超时(%ds)", 20))
				}
				// 此处需return避免协程空跑
				return
			}
		}
	}()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}
	init = true
	// 开启mysql日志
	db = db.Debug()
	global.DB = db

}
