package core

// go build -buildmode=plugin -o=modbusdriver.so main.go
import (
	"plantain/transfer"
)

type rtdbMethod struct {
	collectorName string
	memoryBlock   *MemoryBlock
}

func NewRtdbMethod(
	collectorName string,
	memoryBlock *MemoryBlock,
	alarmTransfer *transfer.AlarmHistoryTranfer,
	historical *transfer.HistoricalTransfer,
) *rtdbMethod {
	return &rtdbMethod{
		collectorName: collectorName,
		memoryBlock:   memoryBlock,
	}
}

func (m *rtdbMethod) Write(pid string, value interface{}) bool {
	if err := m.memoryBlock.WriteInCache(m.collectorName, pid, value); err != nil {
		return false
	}
	//oldVal, found := m.memoryBlock.ReadFromCache(m.collectorName,pid)

	// if found == nil {
	// 	if oldVal != value {
	// 		m.alarm.AlarmJuddge(pid, value)
	// 		m.historical.HistoricalJuddge(pid, value)
	// 	}
	// } else {
	// 	panic("Write Method:内存结构出现错误")
	// }

	// m.cache.Set(pid, value, cache.NoExpiration)
	return true
}

func (m *rtdbMethod) Read(pid string) interface{} {
	if val, err := m.memoryBlock.ReadFromCache(m.collectorName, pid); err == nil {
		return val
	}
	return nil
}
