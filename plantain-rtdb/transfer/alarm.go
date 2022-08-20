package transfer

import (
	"log"
	"plantain/base"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type AlarmHistory struct {
	alarmHistoryChan chan base.AlarmHistoryMessage
	database         string
	token            string
	bucket           string
	url              string
}

func NewAlarmHistory(c *base.AlarmTranferConf) *AlarmHistory {
	return &AlarmHistory{
		alarmHistoryChan: make(chan base.AlarmHistoryMessage, 100),
		database:         c.Database,
		token:            c.Token,
		bucket:           c.Bucket,
		url:              c.Url,
	}
}

func (a *AlarmHistory) Start() {
	log.Println("启动历史报警队列")
	go func() {
		log.Println("开始报警队列接收循环")
		for {
			alarmMessage := <-a.alarmHistoryChan
			go a.writeHistoryAlarmToInfluxdb(alarmMessage)
		}
	}()

}
func (a *AlarmHistory) AddAlarm(am base.AlarmHistoryMessage) {
	log.Printf("将报警 %v 添加到alarmHistoryChan中\n", am.AlarmDes)
	a.alarmHistoryChan <- am
}

func (a *AlarmHistory) writeHistoryAlarmToInfluxdb(am base.AlarmHistoryMessage) {
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
