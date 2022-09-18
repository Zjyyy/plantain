package initiate

import (
	"plantain/base"
	"plantain/collector"
	"plantain/core"
	"plantain/models"
	"plantain/server"
	"plantain/transfer"
)

func ConfigurationAlarmTransfer(conf *base.AlarmTranferConf) *transfer.AlarmHistoryTranfer {
	return transfer.NewAlarmHistoryTranfer(conf)
}

func ConfigurationHistoricalTransfer(conf *base.HistoricalTranferConf) *transfer.HistoricalTransfer {
	return transfer.NewHistoricalTransfer(conf)
}

func ConfigurationMemoryBlockSet(collectorArr *[]models.CollectorWithRtTable) *core.MemoryBlock {
	memoryBlock := *core.NewMemoryBlock()
	memoryBlock.BuildMemoryBlockSet(collectorArr)
	return &memoryBlock
}

func ConfigurationCollector(cp *collector.CollectorParameters) *collector.DriverManager {
	return collector.InitCollector(cp)
}

func ConfigurationHttpServer(port string) {
	server.RouterWeb(port)
}
