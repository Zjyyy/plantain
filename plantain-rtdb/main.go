package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"plantain/base"
	bSqlite "plantain/base/sqlite"
	"plantain/collector"
	"plantain/core"
	"plantain/server"
	"plantain/transfer"
	"syscall"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/**************初始化校验***********************/
	fmt.Println("开始加载配置文件...")
	config, err := base.LoadConfigFromIni("config.ini")
	if err != nil {
		fmt.Printf("加载配置文件失败：%v \n", err)
		return
	}
	//根据配置文件初始化AlarmHistory
	transfer.NewAlarmHistory(config.Alarm)

	//根据配置文件初始化HistoryServe
	//根据配置文件初始化Notice

	fmt.Println("开始项目初始化校验...")
	databaseName := config.System.Database

	_, err = os.Stat(databaseName)
	var db *gorm.DB

	_, err = os.Stat(databaseName)

	db, dbErr := gorm.Open(sqlite.Open(databaseName))
	if dbErr != nil {
		fmt.Printf("打开SQLite连接错误：%v \n", err)
		return
	}

	if err != nil || os.IsNotExist(err) {
		fmt.Println("第一次部署项目，需要创建SQLite")
		db.AutoMigrate(
			&base.RtTable{},
			&base.PDriverInDatabase{},
		)
		bSqlite.CreateMockData(db)
	}
	/**************加载配置库***********************/
	pDriverArr, err := bSqlite.LoadAllDriver(db)
	if err != nil {
		fmt.Printf("加载配置库失败:%v\n", err)
	}

	/**************为配置库实时表建立内存结构***********************/
	cacheSet := core.BuildMemoryStructure(pDriverArr)
	fmt.Println("已在内存中建立了内存表")
	/**************加载驱动插件***********************/
	m := collector.InitCollector(pDriverArr, cacheSet)
	m.Start()
	/**************启动HttpServer***********************/
	server.RouterWeb(":6280")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Println("plantain程序退出")
		return
	}
}
