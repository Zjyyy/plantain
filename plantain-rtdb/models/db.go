package models

import (
	"log"
	"os"
	"plantain/base"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var dbErr error

func InitDb(conf *base.SqliteConf) {

	_, err := os.Stat(conf.Database)
	db, dbErr = gorm.Open(sqlite.Open(conf.Database), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if dbErr != nil {
		log.Printf("打开SQLite连接错误：%v \n", dbErr)
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	if err != nil || os.IsNotExist(err) {
		log.Printf("当前项目下没有Plantain配置库，自动创建样例配置库\n")
		db.AutoMigrate(
			&Collector{},
			&RtTable{},
		)

		CreateMockData()
	}
}
