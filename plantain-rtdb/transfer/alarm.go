package transfer

import (
	"fmt"
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

func NewAlarmHistory(c base.Alarm) *AlarmHistory {
	return &AlarmHistory{
		alarmHistoryChan: make(chan base.AlarmHistoryMessage),
		database:         c.Database,
		token:            c.Token,
		bucket:           c.Bucket,
		url:              c.Url,
	}
}

func (a *AlarmHistory) Start() {
	log.Println("启动历史报警队列")
	alarmMessage := <-a.alarmHistoryChan
	for {
		go a.writeHistoryAlarmToInfluxdb(alarmMessage)
	}
}
func (a *AlarmHistory) AddAlarm(am base.AlarmHistoryMessage) {
	log.Printf("将报警 %v 添加到alarmHistoryChan中\n", am.AlarmDes)
	a.alarmHistoryChan <- am
}

func (a *AlarmHistory) writeHistoryAlarmToInfluxdb(am base.AlarmHistoryMessage) {
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.%v \n", am.Des)
	client := influxdb2.NewClient(a.url, a.token)

	// get non-blocking write client
	writeAPI := client.WriteAPI(a.database, a.bucket)

	// write line protocol
	writeAPI.WriteRecord(
		fmt.Sprintf("%s,unit=string value=%s,valueType=%s,Des=%s,AlarmDes=%s",
			am.Table, am.Value, am.ValueType, am.Des, am.AlarmDes))

	// Flush writes
	writeAPI.Flush()
	client.Close()
}
