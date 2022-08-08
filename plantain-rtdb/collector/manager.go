package collector

import (
	"log"
	"plantain-common/common"
	"plantain/base"
	"plugin"
)

type driverManager struct {
	driverPlugins []DriverPlugin
	rtdbMethods   common.IRTDB
}
type DriverPlugin struct {
	Configure     common.DriverConfigure
	PluginHandler common.IDriver
}

func InitCollector(pDriverArr []base.PDriver, rtdbMethods common.IRTDB) *driverManager {
	driverPlugins := make([]DriverPlugin, len(pDriverArr))

	for index, item := range pDriverArr {
		driverPlugins[index] = DriverPlugin{
			Configure:     createCommonDriverConfigure(item),
			PluginHandler: loadPlugin(item.DriverDllPath),
		}
	}
	return &driverManager{
		driverPlugins,
		rtdbMethods,
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
	handler := d.driverPlugins[index].PluginHandler
	handler.Initialize(
		d.driverPlugins[index].Configure,
		d.rtdbMethods)

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

func loadPlugin(path string) common.IDriver {
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
