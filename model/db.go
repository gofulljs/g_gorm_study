package model

import (
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	// 初始化GORM日志配置
	gCfg := &gorm.Config{
		CreateBatchSize: 100000,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		),
	}

	db, err := gorm.Open(mysql.Open(dsn), gCfg)
	// Error
	if err != nil {
		return errors.WithStack(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return errors.WithStack(err)
	}

	// 设置连接池
	// 空闲
	sqlDB.SetMaxIdleConns(5)
	// 打开
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db
	return nil
}
