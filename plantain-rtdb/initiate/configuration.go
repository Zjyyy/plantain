package initiate

import (
	"plantain/base"
	"plantain/collector"
	"plantain/core"
	"plantain/models"
	"plantain/pipeline"
	"plantain/server"
)

func ConfigurationAlarmPipeline(conf *base.AlarmTranferConf) *pipeline.AlarmPipeline {
	return pipeline.NewAlarmPipeline(conf)
}

func ConfigurationHistoricalPipeline(conf *base.HistoricalTranferConf) *pipeline.HistoricalPipeline {
	return pipeline.NewHistoricalTransfer(conf)
}

func ConfigurationMemoryBlockSet(collectorArr *[]models.CollectorWithRtTable) *core.MemoryBlock {
	memoryBlock := *core.NewMemoryBlock()
	memoryBlock.BuildMemoryBlockSet(collectorArr)
	return &memoryBlock
}

func ConfigurationCollector(cp *collector.CollectorParameters) *collector.CollectorManager {
	return collector.NewCollectorManager(cp)
}

func ConfigurationHttpServer(port string) {
	server.RouterWeb(port)
}
