package core

// go build -buildmode=plugin -o=modbusdriver.so main.go
import (
	"fmt"
	"plantain/models"
	"plantain/pipeline"
)

type rtdbMethod struct {
	collectorName string
	memoryBlock   *MemoryBlock
	alarm         *pipeline.AlarmPipeline
	historical    *pipeline.HistoricalPipeline
}

func NewRtdbMethod(
	collectorName string,
	memoryBlock *MemoryBlock,
	alarm *pipeline.AlarmPipeline,
	historical *pipeline.HistoricalPipeline,
) *rtdbMethod {
	return &rtdbMethod{
		collectorName: collectorName,
		memoryBlock:   memoryBlock,
		alarm:         alarm,
		historical:    historical,
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	oldVal, found := m.memoryBlock.ReadFromCache(m.collectorName, pid)

	if err := m.memoryBlock.WriteInCache(m.collectorName, pid, value); err != nil {
		return false
	}

	if found == nil {
		if oldVal != value {
			m.alarm.AddAlarm(models.AlarmMessage{
				Table:     m.collectorName,
				PID:       pid,
				Des:       "",
				AlarmDes:  "",
				ValueType: "",
				Value:     fmt.Sprintf("%v", value),
			})
			m.historical.AddHistorical(models.HistoricalMessage{
				Table:     m.collectorName,
				PID:       pid,
				Value:     fmt.Sprintf("%v", value),
				ValueType: "",
				Des:       "",
			})
		}
	} else {
		panic("Write Method:内存结构出现错误")
	}

	// m.cache.Set(pid, value, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) interface{} {
	if val, err := m.memoryBlock.ReadFromCache(m.collectorName, pid); err == nil {
		return val
	}
	return nil
}
