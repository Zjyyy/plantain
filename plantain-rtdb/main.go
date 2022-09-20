package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"plantain/collector"
	"plantain/initiate"
	"plantain/models"
	"plantain/server"
	"syscall"
)

func main() {
	go func() {
		for {
			conf, err := initiate.LoadLocalInIConfiguration()
			if err != nil {
				panic(fmt.Sprintf("加载config.ini失败：%v \n", err))
			}

			models.InitDb(&conf.Sqlite)
			collectorWithRtTableSetArr, err := initiate.LoadAllCollectorConfigure(&conf.Sqlite)
			if err != nil {
				log.Fatalln(fmt.Sprintf("从配置库中加载Collector错误:%v \n", err))
			}

			memoryBlock := initiate.ConfigurationMemoryBlockSet(&collectorWithRtTableSetArr)
			println(memoryBlock)
			alarmTransfer := initiate.ConfigurationAlarmTransfer(&conf.AlarmTranfer)
			historicalTranfer := initiate.ConfigurationHistoricalTransfer(&conf.HistoricalTranfer)

			collectorManager := initiate.ConfigurationCollector(&collector.CollectorParameters{
				CollectorArr:       &collectorWithRtTableSetArr,
				MemoryBlock:        memoryBlock,
				AlarmTransfer:      alarmTransfer,
				HistoricalTransfer: historicalTranfer,
			})

			collectorManager.Start()
			select {
			case <-server.RestartChan:
				continue
			}
		}
	}()

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
