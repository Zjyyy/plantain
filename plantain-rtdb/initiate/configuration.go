package initiate

import (
	"plantain/base"
	"plantain/collector"
	"plantain/core"
	"plantain/server"
	"plantain/transfer"

	"github.com/patrickmn/go-cache"
)

func ConfigurationTransferAlarm(conf *base.AlarmTranferConf) *transfer.AlarmHistory {
	return transfer.NewAlarmHistory(conf)
}

func ConfigurationMemoryStructure(driverArr *[]base.PDriver) map[string]*cache.Cache {
	return core.BuildMemoryStructure(driverArr)
}

func ConfigurationCollector(cp *collector.CollectorParameters) *collector.DriverManager {
	return collector.InitCollector(cp)
}

func ConfigurationHttpServer(port string) {
	server.RouterWeb(port)
}
