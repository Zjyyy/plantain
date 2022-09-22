package collector

import (
	"log"
	"plantain-common/common"
	"plantain/core"
	"plantain/models"
	"plantain/pipeline"
	"plugin"
)

type CollectorManager struct {
	*collectorManager
}
type collectorManager struct {
	collectorPlugins []CollectorPlugin
}

type CollectorPlugin struct {
	Configure     common.DriverConfigure
	RTDBHandler   common.IRTDB
	PluginHandler common.IDriver
}

type CollectorParameters struct {
	CollectorArr       *[]models.CollectorWithRtTable
	MemoryBlock        *core.MemoryBlock
	AlarmTransfer      *pipeline.AlarmPipeline
	HistoricalPipeline *pipeline.HistoricalPipeline
}

func NewCollectorManager(collectorParameters *CollectorParameters) *CollectorManager {
	return &CollectorManager{
		newCollectorManager(collectorParameters),
	}
}

func newCollectorManager(cp *CollectorParameters) *collectorManager {
	collectorPlugins := make([]CollectorPlugin, len(*cp.CollectorArr))

	for index, collector := range *cp.CollectorArr {

		rtdbHandler := core.NewRtdbMethod(
			collector.CollectorName,
			cp.MemoryBlock,
			cp.AlarmTransfer,
			cp.HistoricalPipeline,
		)

		collectorPlugins[index] = CollectorPlugin{
			RTDBHandler:   rtdbHandler,
			Configure:     createDriverConfigure(collector),
			PluginHandler: loadDriverPluginHandler(collector.DllPath),
		}
	}
	return &collectorManager{
		collectorPlugins,
	}
}

func (d *collectorManager) Start() {
	for index, item := range d.collectorPlugins {
		log.Printf("开始执行< %v >驱动",
			item.Configure.DriverName)
		go d.taskWork(index)
	}
}

func (d *collectorManager) taskWork(index int) {
	handler := d.collectorPlugins[index].PluginHandler
	handler.Initialize(
		d.collectorPlugins[index].Configure,
		d.collectorPlugins[index].RTDBHandler)
	for {
		handler.Do()
	}
}

func createDriverConfigure(collector models.CollectorWithRtTable) common.DriverConfigure {
	rtPoint := make(map[string]string, len(collector.RtTableSet))

	for _, item := range collector.RtTableSet {
		if _, ok := rtPoint[item.PID]; !ok {
			rtPoint[item.PID] = item.Address
		}
	}

	return common.DriverConfigure{
		Id:         collector.Id,
		DriverName: collector.CollectorName,
		ConnStr:    collector.ConnStr,
		RtPoint:    rtPoint,
	}
}

func loadDriverPluginHandler(path string) common.IDriver {
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
