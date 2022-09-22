package core

import (
	"errors"
	"fmt"
	"plantain/models"

	"github.com/patrickmn/go-cache"
)

// 供plantain-RTDB操作
var MemoryBlockHandler *MemoryBlock

type MemoryBlock struct {
	*memoryBlock
}
type memoryBlock struct {
	collectorCacheMap map[string]*cache.Cache
}

func NewMemoryBlock() *MemoryBlock {
	return &MemoryBlock{
		newMemoryBlock(),
	}
}
func newMemoryBlock() *memoryBlock {
	return &memoryBlock{
		collectorCacheMap: make(map[string]*cache.Cache),
	}
}

// 为每一个Collector开辟内存空间
func (m *memoryBlock) BuildMemoryBlockSet(collectorArr *[]models.CollectorWithRtTable) {

	for _, collector := range *collectorArr {
		_, ok := m.collectorCacheMap[collector.CollectorName]
		if ok {
			panic("存在两个重复驱动名，请检查配置，DriverName：" + collector.CollectorName)
		}

		m.collectorCacheMap[collector.CollectorName] = cache.New(0, 0)

		for _, rtItem := range collector.RtTableSet {
			_, found := m.collectorCacheMap[collector.CollectorName].Get(rtItem.PID)
			if found {
				panic("PID已存在，请保证PID在RTDB中是唯一的，PID:" + rtItem.PID)
			}
			m.collectorCacheMap[collector.CollectorName].Set(rtItem.PID, rtItem.Value, cache.NoExpiration)
		}
	}
}

func (m memoryBlock) GetCache(collectorName string) (*cache.Cache, error) {
	if cache, ok := m.collectorCacheMap[collectorName]; ok {
		return cache, nil
	} else {
		return nil, errors.New(fmt.Sprintf("无法找到 %v 的内存块 \n", collectorName))
	}
}

func (m *memoryBlock) ReadFromCache(collectorName string, pid string) (interface{}, error) {
	if cache, err := m.GetCache(collectorName); err == nil {
		if val, found := cache.Get(pid); found {
			return val, nil
		} else {
			return nil, errors.New(fmt.Sprintf("无法在 %v 的内存块中找到PID: %v \n", collectorName, pid))
		}
	} else {
		return nil, err
	}
}

func (m *memoryBlock) WriteInCache(collectorName string, pid string, value interface{}) error {
	if cache, ok := m.collectorCacheMap[collectorName]; ok {
		if _, found := cache.Get(pid); found {
			cache.Set(pid, value, -1)
			return nil
		} else {
			return errors.New(fmt.Sprintf("无法在 %v 的内存块中找到PID: %v \n", collectorName, pid))
		}
	} else {
		return errors.New(fmt.Sprintf("无法找到 %v 的内存块 \n", collectorName))
	}
}

func (m *memoryBlock) CountMemoryBlockSet() int {
	return len(m.collectorCacheMap)
}
