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
	conf, err := initiate.LoadLocalInIConfiguration()
	if err != nil {
		panic(fmt.Sprintf("加载config.ini失败：%v \n", err))
	}

	driverArr, err := initiate.LoadSQLiteConfiguration(&conf.Sqlite)
	if err != nil {
		panic(fmt.Sprintf("打开SQLite连接错误：%v \n", err))
	}

	alarmTransfer := initiate.ConfigurationTransferAlarm(&conf.AlarmTranfer)

	cacheMap := initiate.ConfigurationMemoryStructure(&driverArr)

	collector := initiate.ConfigurationCollector(&collector.CollectorParameters{
		DriverArr:     &driverArr,
		CacheSet:      &cacheMap,
		AlarmTransfer: alarmTransfer,
	})
	collector.Start()

	initiate.ConfigurationHttpServer(":6280")

	//退出程序
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Println("plantain程序退出")
		return
	}
}
