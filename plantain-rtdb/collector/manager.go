package collector

import (
	"log"
	"plantain-common/common"
	"plantain/base"
	"plantain/core"
	"plugin"

	"github.com/patrickmn/go-cache"
)

type driverManager struct {
	driverPlugins []DriverPlugin
}
type DriverPlugin struct {
	Configure     common.DriverConfigure
	RTDBHandler   common.IRTDB
	DriverHandler common.IDriver
}

func InitCollector(pDriverArr []base.PDriver, cacheSet map[string]*cache.Cache) *driverManager {
	driverPlugins := make([]DriverPlugin, len(pDriverArr))

	for index, pDriver := range pDriverArr {
		driverPlugins[index] = DriverPlugin{
			RTDBHandler:   core.NewRtdbMethod(pDriver, cacheSet[pDriver.DriverName]),
			Configure:     createCommonDriverConfigure(pDriver),
			DriverHandler: loadDriverPlugin(pDriver.DriverDllPath),
		}
	}
	return &driverManager{
		driverPlugins,
	}
}

func (d *driverManager) Start() {
	for index, item := range d.driverPlugins {
		log.Printf("开始执行< %v >驱动",
			item.Configure.DriverName)
		go d.taskWork(index)
	}
}

func (d *driverManager) taskWork(index int) {
	handler := d.driverPlugins[index].DriverHandler
	handler.Initialize(
		d.driverPlugins[index].Configure,
		d.driverPlugins[index].RTDBHandler)
	for {
		handler.Do()
	}
}

func createCommonDriverConfigure(conf base.PDriver) common.DriverConfigure {
	rtPoint := make(map[string]string, len(conf.RtTable))

	for _, item := range conf.RtTable {
		if _, ok := rtPoint[item.PID]; !ok {
			rtPoint[item.PID] = item.Address
		}
	}

	return common.DriverConfigure{
		Id:         conf.Id,
		DriverName: conf.DriverName,
		ConnStr:    conf.ConnStr,
		RtPoint:    rtPoint,
	}
}

func loadDriverPlugin(path string) common.IDriver {
	plug, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}

	noType, err := plug.Lookup("Driver")
	if err != nil {
		panic(err)
	}

	p, ok := noType.(common.IDriver)
	if !ok {
		panic("load plugin error")
	}
	return p
}
