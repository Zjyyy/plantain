package monitor

import (
	"plantain/base"
	"strconv"
)

type MonitorAlarm struct {
	*monitorAlarm
}

type monitorAlarm struct {
	alarmConfMap AlarmConfMap
}

type alarmConfItem struct {
	Des       string
	ValueType string
	LimitUp   string
	LimitDown string
	Level     uint
	AlarmDes  string
}
type AlarmConfMap map[string]alarmConfItem

func NewMonitorAlarm(pDriver *base.PDriver) *MonitorAlarm {
	return &MonitorAlarm{
		monitorAlarm: newMonitorAlarm(pDriver),
	}
}

func newMonitorAlarm(pDriver *base.PDriver) *monitorAlarm {
	return &monitorAlarm{
		alarmConfMap: parseForAlarm(pDriver),
	}
}

func (m *monitorAlarm) AlarmHandler(pid string, val interface{}) {
	item := m.alarmConfMap[pid]
	if item.ValueType == "int" {
		standardLimitUp, _ := strconv.Atoi(item.LimitUp)
		standardLimitDown, _ := strconv.Atoi(item.LimitDown)

		if val.(int) > standardLimitUp || val.(int) < standardLimitDown {
			//触发报警
		}
	} else if item.ValueType == "float" {
		standardLimitUp, _ := strconv.ParseFloat(item.LimitUp, 64)
		standardLimitDown, _ := strconv.ParseFloat(item.LimitDown, 64)

		if val.(float64) > standardLimitUp || val.(float64) < standardLimitDown {
			//触发报警
		}
	} else if item.ValueType == "boolen" {
		standardLimitUp, _ := strconv.ParseBool(item.LimitUp)

		if val.(bool) == standardLimitUp {
			//触发报警
		}
	}
}

func parseForAlarm(pDriver *base.PDriver) AlarmConfMap {
	result := make(map[string]alarmConfItem)
	for _, item := range pDriver.RtTable {
		result[item.PID] = alarmConfItem{
			Des:       item.Des,
			ValueType: item.ValueType,
			LimitUp:   item.LimitUp,
			LimitDown: item.LimitDown,
			Level:     1,
			AlarmDes:  item.AlarmDes,
		}
	}
	return result
}
