package monitor

import (
	"plantain/models"
	"plantain/pipeline"
	"strconv"
)

type MonitorAlarm struct {
	*monitorAlarm
}

type monitorAlarm struct {
	name          string
	alarmPipeline *pipeline.AlarmPipeline
	alarmConfMap  AlarmConfMap
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

func NewMonitorAlarm(
	collector *models.CollectorWithRtTable,
	alarmTransfer *pipeline.AlarmPipeline,
) *MonitorAlarm {
	return &MonitorAlarm{
		monitorAlarm: newMonitorAlarm(collector, alarmTransfer),
	}
}

func newMonitorAlarm(
	collector *models.CollectorWithRtTable,
	alarmPipeline *pipeline.AlarmPipeline,
) *monitorAlarm {
	return &monitorAlarm{
		name:          collector.CollectorName,
		alarmPipeline: alarmPipeline,
		alarmConfMap:  parseForAlarm(collector),
	}
}

func (m *monitorAlarm) AlarmJuddge(pid string, val interface{}) {
	item := m.alarmConfMap[pid]
	var isAlarm bool = false
	var alarmValue string = ""
	switch item.ValueType {
	case "int":
		isAlarm, alarmValue = m.juddgeInt(&item, val)
		break
	case "float":
		isAlarm, alarmValue = m.juddgeFloat(&item, val)
		break
	case "uint16":
		isAlarm, alarmValue = m.juddgeUint16(&item, val)
		break
	case "bool":
		isAlarm, alarmValue = m.juddgeBool(&item, val)
		break
	}
	if isAlarm {
		m.alarmPipeline.AddAlarm(models.AlarmMessage{
			Table:     m.name,
			PID:       pid,
			Des:       item.Des,
			AlarmDes:  item.AlarmDes,
			ValueType: item.ValueType,
			Value:     alarmValue,
		})
	}
}

func parseForAlarm(collector *models.CollectorWithRtTable) AlarmConfMap {
	result := make(map[string]alarmConfItem)
	for _, item := range collector.RtTableSet {
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

func (m monitorAlarm) juddgeInt(item *alarmConfItem, val interface{}) (bool, string) {
	standardLimitUp, _ := strconv.Atoi(item.LimitUp)
	standardLimitDown, _ := strconv.Atoi(item.LimitDown)

	if val.(int) > standardLimitUp || val.(int) < standardLimitDown {
		return true, val.(string)
	}
	return false, ""
}
func (m monitorAlarm) juddgeFloat(item *alarmConfItem, val interface{}) (bool, string) {
	standardLimitUp, _ := strconv.ParseFloat(item.LimitUp, 64)
	standardLimitDown, _ := strconv.ParseFloat(item.LimitDown, 64)

	if val.(float64) > standardLimitUp || val.(float64) < standardLimitDown {
		return true, val.(string)
	}
	return false, ""
}
func (m monitorAlarm) juddgeBool(item *alarmConfItem, val interface{}) (bool, string) {
	standardLimitUp, _ := strconv.ParseBool(item.LimitUp)

	if val.(bool) == standardLimitUp {
		return true, val.(string)
	}
	return false, ""
}
func (m monitorAlarm) juddgeUint16(item *alarmConfItem, val interface{}) (bool, string) {
	standardLimitUp, _ := strconv.Atoi(item.LimitUp)
	standardLimitDown, _ := strconv.Atoi(item.LimitDown)

	if val.(uint16) > uint16(standardLimitUp) || val.(uint16) < uint16(standardLimitDown) {
		return true, string(val.(uint16))
	}
	return false, ""
}
