package transfer

import (
	"fmt"
	"log"
	"plantain/base"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type AlarmHistory struct {
	database string
	token    string
	bucket   string
	url      string
}

func NewAlarmHistory(c base.Alarm) *AlarmHistory {
	return &AlarmHistory{
		database: c.Database,
		token:    c.Token,
		bucket:   c.Bucket,
		url:      c.Url,
	}
}

func (a *AlarmHistory) Start() {
	log.Println("启动历史报警队列")
}

func (a *AlarmHistory) AddAlarm(tableName string, value string,
	valueType string, Des string, AlarmDes string) {

	client := influxdb2.NewClient(a.url, a.token)

	// get non-blocking write client
	writeAPI := client.WriteAPI(a.database, a.bucket)

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("%s,unit=string value=%s,valueType=%s,Des=%s,AlarmDes=%s", tableName, value, valueType, Des, AlarmDes))

	// Flush writes
	writeAPI.Flush()
	client.Close()
}
