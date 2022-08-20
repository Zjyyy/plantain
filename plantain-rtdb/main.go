package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"plantain/collector"
	"plantain/initiate"
	"syscall"
)

func main() {
	/**************初始化校验***********************/
	conf, err := initiate.LoadLocalInIConfiguration()
	if err != nil {
		panic(fmt.Sprintf("加载config.ini失败：%v \n", err))
	}

	driverArr, err := initiate.LoadSQLiteConfiguration(&conf.Sqlite)
	if err != nil {
		panic(fmt.Sprintf("打开SQLite连接错误：%v \n", err))
	}

	//根据配置文件初始化AlarmHistory
	alarmTransfer := initiate.ConfigurationTransferAlarm(&conf.AlarmTranfer)

	/**************为配置库实时表建立内存结构***********************/
	cacheMap := initiate.ConfigurationMemoryStructure(&driverArr)
	/**************加载驱动插件***********************/
	m := initiate.ConfigurationCollector(&collector.CollectorParameters{
		DriverArr:     &driverArr,
		CacheSet:      &cacheMap,
		AlarmTransfer: alarmTransfer,
	})
	m.Start()
	/**************启动HttpServer***********************/
	initiate.ConfigurationHttpServer(":6280")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Println("plantain程序退出")
		return
	}
}
