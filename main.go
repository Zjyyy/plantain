package main

import (
	"fmt"
	"os"
	"plantain/base"
	bSqlite "plantain/base/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/**************初始化校验***********************/
	fmt.Println("开始加载配置文件...")
	config, err := base.LoadConfigFromIni("config.ini")
	if err == nil {
		fmt.Printf("加载配置文件失败：%v \n", err)
	}
	fmt.Println("开始项目初始化校验...")
	databaseName := config.System.Database

	_, err = os.Stat(databaseName)
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(databaseName))
	if err != nil {
		fmt.Printf("打开SQLite连接错误：%v \n", err)
		return
	}

	_, err = os.Stat(databaseName)
	if err != nil || os.IsNotExist(err) {
		fmt.Println("第一次部署项目，需要创建SQLite")
		db.AutoMigrate(
			&base.RtTable{},
			&base.PDriverInDatabase{},
		)
		bSqlite.CreateMockData(db)
	}
	/**************加载配置库***********************/

	/**************加载驱动插件***********************/
	/**************为配置库实时表建立内存结构***********************/
	/**************启动HttpServer***********************/
}
