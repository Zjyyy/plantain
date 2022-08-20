package initiate

import (
	"plantain/base"
	"plantain/collector"
	"plantain/core"
	"plantain/server"
	"plantain/transfer"

	"github.com/patrickmn/go-cache"
)

func ConfigurationAlarmTransfer(conf *base.AlarmTranferConf) *transfer.AlarmHistoryTranfer {
	return transfer.NewAlarmHistoryTranfer(conf)
}

func ConfigurationHistoricalTransfer(conf *base.HistoricalTranferConf) *transfer.HistoricalTransfer {
	return transfer.NewHistoricalTransfer(conf)
}

func ConfigurationMemoryBlockSet(driverArr *[]base.PDriver) map[string]*cache.Cache {
	return core.BuildMemoryBlockSet(driverArr)
}

func ConfigurationCollector(cp *collector.CollectorParameters) *collector.DriverManager {
	return collector.InitCollector(cp)
}

func ConfigurationHttpServer(port string) {
	server.RouterWeb(port)
}
