package monitor

import (
	"log"
	"plantain/base"
)

type MonitorHistorical struct {
	monitorHistorical *monitorHistorical
}

type monitorHistorical struct {
	historicalConfMap HistoricalConfMap
}

type HistoricalConfMap map[string]bool

func NewMonitorHistorical(pDriver *base.PDriver) *MonitorHistorical {
	return &MonitorHistorical{
		newMonitorHistorical(pDriver),
	}
}

func newMonitorHistorical(pDriver *base.PDriver) *monitorHistorical {
	return &monitorHistorical{
		parseForHistorical(pDriver),
	}
}

func (m *monitorHistorical) HistoricalJuddge(pid string, val interface{}) {
	if m.historicalConfMap[pid] == true {
		//将变动操作存放到历史库
		log.Printf("放到历史库中值，%v", pid)
	}
}

func parseForHistorical(pDriver *base.PDriver) HistoricalConfMap {
	result := make(HistoricalConfMap)
	for _, item := range pDriver.RtTable {
		result[item.PID] = item.IsHistorical
	}
	return result
}
