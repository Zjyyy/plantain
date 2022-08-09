package core

import (
	"plantain/base"

	"github.com/patrickmn/go-cache"
)

func BuildMemoryStructure(driverArr []base.PDriver) map[string]*cache.Cache {
	cacheSet := make(map[string]*cache.Cache)

	for _, driver := range driverArr {
		_, ok := cacheSet[driver.DriverName]
		if ok {
			panic("存在两个重复驱动名，请检查配置，DriverName：" + driver.DriverName)
		}

		cacheSet[driver.DriverName] = cache.New(0, 0)

		for _, rtItem := range driver.RtTable {
			_, found := cacheSet[driver.DriverName].Get(rtItem.PID)
			if found {
				panic("PID已存在，请保证PID在RTDB中是唯一的，PID:" + rtItem.PID)
			}
			cacheSet[driver.DriverName].Set(rtItem.PID, rtItem.Value, cache.NoExpiration)
		}
	}
	return cacheSet
}
