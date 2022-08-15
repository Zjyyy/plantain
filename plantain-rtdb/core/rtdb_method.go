package core

import (
	"plantain/base"
	"plantain/core/monitor"

	"github.com/patrickmn/go-cache"
)

type rtdbMethod struct {
	cache *cache.Cache
	alarm *monitor.MonitorAlarm
}

func NewRtdbMethod(pDriver base.PDriver, cache *cache.Cache) *rtdbMethod {
	return &rtdbMethod{
		cache: cache,
		alarm: monitor.NewMonitorAlarm(&pDriver),
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	oldVal, found := m.cache.Get(pid)

	if found {
		if oldVal != value {
			m.alarm.AlarmHandler(pid, value)
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
