package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"plantain/initiate"
	"plantain/models"
	"syscall"
)

func main() {
	conf, err := initiate.LoadLocalInIConfiguration()
	if err != nil {
		panic(fmt.Sprintf("加载config.ini失败：%v \n", err))
	}

	models.InitDb(&conf.Sqlite)
	collectorWithRtTableSetArr, err := initiate.LoadAllCollectorWithRTTableSet(&conf.Sqlite)
	if err != nil {
		panic(fmt.Sprintf("从配置库中加载Collector错误:%v \n", err))
	}

	memoryBlock := initiate.ConfigurationMemoryBlockSet(&collectorWithRtTableSetArr)
	println(memoryBlock)

	// alarmTransfer := initiate.ConfigurationAlarmTransfer(&conf.AlarmTranfer)
	// historicalTranfer := initiate.ConfigurationHistoricalTransfer(&conf.HistoricalTranfer)

	// collector := initiate.ConfigurationCollector(&collector.CollectorParameters{
	// 	DriverArr:          &driverArr,
	// 	MemoryBlockSet:     &memoryBlockSet,
	// 	AlarmTransfer:      alarmTransfer,
	// 	HistoricalTransfer: historicalTranfer,
	// })
	// collector.Start()

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
