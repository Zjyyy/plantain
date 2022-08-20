package base

import (
	"github.com/go-ini/ini"
)

type SqliteConf struct {
	Database            string `ini:"database"`
	DriverListTableName string `int:"driver-list-table-name"`
	DriverTableName     string `ini:"driver-table-name"`
}
type HistoricalTranferConf struct {
	Url      string `ini:"historical-tranfer-url"`
	Token    string `ini:"historical-tranfer-token"`
	Database string `ini:"historical-tranfer-database"`
	Bucket   string `ini:"historical-tranfer-bucket"`
}
type AlarmTranferConf struct {
	Url      string `ini:"alarm-tranfer-url"`
	Token    string `ini:"alarm-tranfer-token"`
	Database string `ini:"alarm-tranfer-database"`
	Bucket   string `ini:"alarm-tranfer-bucket"`
}
type Config struct {
	Sqlite            SqliteConf            `ini:"sqlite"`
	AlarmTranfer      AlarmTranferConf      `ini:"alarm-tranfer"`
	HistoricalTranfer HistoricalTranferConf `ini:"historical-tranfer"`
}

func LoadConfigFromIni(fileName string) (Config, error) {
	CONFIG := new(Config)
	err := ini.MapTo(CONFIG, fileName)
	if err != nil {
		return *CONFIG, nil
	}
	return *CONFIG, err
}
