package monitor

import (
	"fmt"
	"plantain/base"
	"plantain/transfer"
)

type MonitorHistorical struct {
	*monitorHistorical
}

type monitorHistorical struct {
	table              string
	historicalTransfer *transfer.HistoricalTransfer
	historicalConfMap  HistoricalConfMap
}

type HistoricalConfMap map[string]historicalConfItem
type historicalConfItem struct {
	IsHistorical bool
	Des          string
	ValueType    string
}

func NewMonitorHistorical(pDriver *base.PDriver, historicalTransfer *transfer.HistoricalTransfer) *MonitorHistorical {
	return &MonitorHistorical{
		newMonitorHistorical(pDriver, historicalTransfer),
	}
}

func newMonitorHistorical(pDriver *base.PDriver, historicalTransfer *transfer.HistoricalTransfer) *monitorHistorical {
	return &monitorHistorical{
		table:              pDriver.DriverName,
		historicalTransfer: historicalTransfer,
		historicalConfMap:  parseForHistorical(pDriver),
	}
}

func (m *monitorHistorical) HistoricalJuddge(pid string, val interface{}) {
	item := m.historicalConfMap[pid]
	if item.IsHistorical {
		//将变动操作存放到历史库
		if item.ValueType == "int" {
			m.historicalTransfer.AddHistorical(base.HistoricalMessage{
				Table:     m.table,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprintf("%d", val.(int)),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "float" {
			m.historicalTransfer.AddHistorical(base.HistoricalMessage{
				Table:     m.table,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprintf("%f", val.(float64)),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "string" {
			m.historicalTransfer.AddHistorical(base.HistoricalMessage{
				Table:     m.table,
				PID:       pid,
				Des:       item.Des,
				Value:     val.(string),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "boolen" {
			m.historicalTransfer.AddHistorical(base.HistoricalMessage{
				Table:     m.table,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprint(val.(bool)),
				ValueType: item.ValueType,
			})
		}

	}
}

func parseForHistorical(pDriver *base.PDriver) HistoricalConfMap {
	result := make(HistoricalConfMap)
	for _, item := range pDriver.RtTable {
		result[item.PID] = historicalConfItem{
			IsHistorical: item.IsHistorical,
			Des:          item.Des,
			ValueType:    item.ValueType,
		}
	}
	return result
}
