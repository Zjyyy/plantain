package pipeline

import (
	"log"
	"plantain/base"
	"plantain/models"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type HistoricalPipeline struct {
	historicalChan chan models.HistoricalMessage
	database       string
	token          string
	bucket         string
	url            string
}

func NewHistoricalTransfer(c *base.HistoricalTranferConf) *HistoricalPipeline {
	return &HistoricalPipeline{
		historicalChan: make(chan models.HistoricalMessage, 100),
		database:       c.Database,
		url:            c.Url,
		bucket:         c.Bucket,
		token:          c.Token,
	}
}

func (self *HistoricalPipeline) Start() {
	log.Println("启动历史归档队列")
	go func() {
		log.Println("开始历史归档队列接收循环")
		for {
			historicalMessage := <-self.historicalChan
			go self.writeHistoricalMessageToInfluxdb(historicalMessage)
		}
	}()
}

func (self *HistoricalPipeline) AddHistorical(hm models.HistoricalMessage) {
	log.Printf("将 %v 点的值添加到HistoricalChan中\n", hm.PID)
	self.historicalChan <- hm
}
func (self *HistoricalPipeline) writeHistoricalMessageToInfluxdb(hm models.HistoricalMessage) {
	client := influxdb2.NewClient(self.url, self.token)

	// get non-blocking write client
	writeAPI := client.WriteAPI(self.database, self.bucket)

	// write line protocol
	p := influxdb2.NewPointWithMeasurement(hm.Table).
		AddTag("pid", hm.PID).
		AddField("des", hm.Des).
		AddField("value", hm.Value).
		AddField("valueType", hm.ValueType)

	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
	client.Close()
}
