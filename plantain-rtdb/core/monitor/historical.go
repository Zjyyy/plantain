package monitor

import (
	"fmt"
	"plantain/models"
	"plantain/pipeline"
)

type MonitorHistorical struct {
	*monitorHistorical
}

type monitorHistorical struct {
	name               string
	historicalPipeline *pipeline.HistoricalPipeline
	historicalConfMap  HistoricalConfMap
}

type HistoricalConfMap map[string]historicalConfItem
type historicalConfItem struct {
	IsHistorical bool
	Des          string
	ValueType    string
}

func NewMonitorHistorical(
	collector *models.CollectorWithRtTable,
	historicalTransfer *pipeline.HistoricalPipeline,
) *MonitorHistorical {
	return &MonitorHistorical{
		newMonitorHistorical(collector, historicalTransfer),
	}
}

func newMonitorHistorical(collector *models.CollectorWithRtTable, historicalPipeline *pipeline.HistoricalPipeline) *monitorHistorical {
	return &monitorHistorical{
		name:               collector.CollectorName,
		historicalPipeline: historicalPipeline,
		historicalConfMap:  parseForHistorical(collector),
	}
}

func (m *monitorHistorical) HistoricalJuddge(pid string, val interface{}) {
	item := m.historicalConfMap[pid]
	if item.IsHistorical {
		//将变动操作存放到历史库
		if item.ValueType == "int" {
			m.historicalPipeline.AddHistorical(models.HistoricalMessage{
				Table:     m.name,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprintf("%d", val.(int)),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "float" {
			m.historicalPipeline.AddHistorical(models.HistoricalMessage{
				Table:     m.name,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprintf("%f", val.(float64)),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "string" {
			m.historicalPipeline.AddHistorical(models.HistoricalMessage{
				Table:     m.name,
				PID:       pid,
				Des:       item.Des,
				Value:     val.(string),
				ValueType: item.ValueType,
			})
		} else if item.ValueType == "boolen" {
			m.historicalPipeline.AddHistorical(models.HistoricalMessage{
				Table:     m.name,
				PID:       pid,
				Des:       item.Des,
				Value:     fmt.Sprint(val.(bool)),
				ValueType: item.ValueType,
			})
		}

	}
}

func parseForHistorical(collector *models.CollectorWithRtTable) HistoricalConfMap {
	result := make(HistoricalConfMap)
	for _, item := range collector.RtTableSet {
		result[item.PID] = historicalConfItem{
			IsHistorical: item.IsHistorical,
			Des:          item.Des,
			ValueType:    item.ValueType,
		}
	}
	return result
}
