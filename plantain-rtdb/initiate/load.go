package initiate

import (
	"log"
	"os"
	"plantain/base"
	plantainSqlite "plantain/base/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadLocalInIConfiguration() (base.Config, error) {
	config, err := base.LoadConfigFromIni("config.ini")
	if err != nil {
		log.Printf("加载config.ini失败：%v \n", err)
		return base.Config{}, err
	}
	return config, nil
}

func LoadSQLiteConfiguration(conf *base.SqliteConf) ([]base.PDriver, error) {
	var db *gorm.DB
	_, err := os.Stat(conf.Database)

	db, dbErr := gorm.Open(sqlite.Open(conf.Database))
	if dbErr != nil {
		log.Printf("打开SQLite连接错误：%v \n", dbErr)
		return []base.PDriver{}, dbErr
	}

	if err != nil || os.IsNotExist(err) {
		log.Printf("当前项目下没有Plantain配置库，自动创建样例配置库\n")
		db.AutoMigrate(
			&base.RtTable{},
			&base.PDriverInDatabase{},
		)
		plantainSqlite.CreateMockData(db)
	}

	pDriverArr, err := plantainSqlite.LoadAllDriver(db)
	if err != nil {
		log.Printf("加载配置库失败:%v \n", err)
		return []base.PDriver{}, err
	}
	return pDriverArr, nil
}
