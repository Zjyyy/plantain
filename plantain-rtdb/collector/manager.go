package collector

import (
	"log"
	"plantain-common/common"
	"plantain/base"
	"plantain/core"
	"plantain/transfer"
	"plugin"

	"github.com/patrickmn/go-cache"
)

type DriverManager struct {
	*driverManager
}
type driverManager struct {
	driverPlugins []DriverPlugin
}
type DriverPlugin struct {
	Configure     common.DriverConfigure
	RTDBHandler   common.IRTDB
	DriverHandler common.IDriver
}
type CollectorParameters struct {
	DriverArr     *[]base.PDriver
	CacheSet      *map[string]*cache.Cache
	AlarmTransfer *transfer.AlarmHistoryTranfer
}

func InitCollector(collectorParameters *CollectorParameters) *DriverManager {
	return &DriverManager{
		newDriverManager(collectorParameters),
	}
}

func newDriverManager(cp *CollectorParameters) *driverManager {
	driverPlugins := make([]DriverPlugin, len(*cp.DriverArr))

	for index, pDriver := range *cp.DriverArr {
		driverPlugins[index] = DriverPlugin{
			RTDBHandler:   core.NewRtdbMethod(pDriver, (*cp.CacheSet)[pDriver.DriverName], cp.AlarmTransfer),
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
