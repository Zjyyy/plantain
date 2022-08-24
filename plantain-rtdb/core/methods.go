package core

// go build -buildmode=plugin -o=modbusdriver.so main.go
import (
	"plantain/base"
	"plantain/core/monitor"
	"plantain/transfer"

	"github.com/patrickmn/go-cache"
)

type rtdbMethod struct {
	cache      *cache.Cache
	alarm      *monitor.MonitorAlarm
	historical *monitor.MonitorHistorical
}

func NewRtdbMethod(
	pDriver *base.PDriver,
	cache *cache.Cache,
	alarmTransfer *transfer.AlarmHistoryTranfer,
	historical *transfer.HistoricalTransfer,
) *rtdbMethod {
	return &rtdbMethod{
		cache:      cache,
		alarm:      monitor.NewMonitorAlarm(pDriver, alarmTransfer),
		historical: monitor.NewMonitorHistorical(pDriver, historical),
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	oldVal, found := m.cache.Get(pid)

	if found {
		if oldVal != value {
			//m.alarm.AlarmJuddge(pid, value)
			//m.historical.HistoricalJuddge(pid, value)
		}
	} else {
		panic("Write Method:内存结构出现错误")
	}

	m.cache.Set(pid, value, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) interface{} {
	val, found := m.cache.Get(pid)
	if found {
		return val
	}
	return ""
}
