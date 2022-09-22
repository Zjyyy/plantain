package pipeline

import (
	"log"
	"plantain/base"
	"plantain/models"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type AlarmPipeline struct {
	alarmChan chan models.AlarmMessage
	database  string
	token     string
	bucket    string
	url       string
}

func NewAlarmPipeline(c *base.AlarmTranferConf) *AlarmPipeline {
	return &AlarmPipeline{
		alarmChan: make(chan models.AlarmMessage, 100),
		database:  c.Database,
		token:     c.Token,
		bucket:    c.Bucket,
		url:       c.Url,
	}
}

func (a *AlarmPipeline) Start() {
	log.Println("启动历史报警队列")
	go func() {
		log.Println("开始报警队列接收循环")
		for {
			alarmMessage := <-a.alarmChan
			go a.writeAlarmToInfluxdb(alarmMessage)
		}
	}()

}
func (a *AlarmPipeline) AddAlarm(am models.AlarmMessage) {
	log.Printf("将报警 %v 添加到alarmChan中\n", am.AlarmDes)
	a.alarmChan <- am
}

func (a *AlarmPipeline) writeAlarmToInfluxdb(am models.AlarmMessage) {
	client := influxdb2.NewClient(a.url, a.token)

	// get non-blocking write client
	writeAPI := client.WriteAPI(a.database, a.bucket)

	// write line protocol
	p := influxdb2.NewPointWithMeasurement(am.Table).
		AddTag("pid", am.PID).
		AddField("des", am.Des).
		AddField("alarmDes", am.AlarmDes).
		AddField("valueType", am.ValueType).
		AddField("value", am.Value)
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
	client.Close()
}
