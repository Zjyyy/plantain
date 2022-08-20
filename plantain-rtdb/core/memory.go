package core

import (
	"plantain/base"

	"github.com/patrickmn/go-cache"
)

// 一个驱动一个内存map
func BuildMemoryBlockSet(driverArr *[]base.PDriver) map[string]*cache.Cache {
	memoryBlockSet := make(map[string]*cache.Cache)

	for _, driver := range *driverArr {
		_, ok := memoryBlockSet[driver.DriverName]
		if ok {
			panic("存在两个重复驱动名，请检查配置，DriverName：" + driver.DriverName)
		}

		memoryBlockSet[driver.DriverName] = cache.New(0, 0)

		for _, rtItem := range driver.RtTable {
			_, found := memoryBlockSet[driver.DriverName].Get(rtItem.PID)
			if found {
				panic("PID已存在，请保证PID在RTDB中是唯一的，PID:" + rtItem.PID)
			}
			memoryBlockSet[driver.DriverName].Set(rtItem.PID, rtItem.Value, cache.NoExpiration)
		}
	}
	return memoryBlockSet
}
