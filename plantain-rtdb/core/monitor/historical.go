package monitor

import "plantain/base"

type MonitorHistorical struct {
	historicalConfMap HistoricalConfMap
}

type HistoricalConfMap map[string]bool

func NewMonitorHistorical(pDriver *base.PDriver) *MonitorHistorical {
	return &MonitorHistorical{
		parseForHistorical(pDriver),
	}
}

func (m *MonitorHistorical) HistoricalHandler(pid string) {
	if m.historicalConfMap[pid] == true {
		//将变动操作存放到历史库
	}
}

func parseForHistorical(pDriver *base.PDriver) HistoricalConfMap {
	result := make(HistoricalConfMap)
	for _, item := range pDriver.RtTable {
		result[item.PID] = item.IsHistorical
	}
	return result
}
