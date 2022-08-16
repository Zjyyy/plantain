package base

import (
	"github.com/go-ini/ini"
)

type System struct {
	Database string `ini:"database"`
}
type Alarm struct {
	Url      string `ini:"url"`
	Token    string `ini:"token"`
	Database string `ini:"database"`
	Bucket   string `ini:"bucket"`
}
type Config struct {
	System System `ini:"system"`
	Alarm  Alarm  `ini:"alarm"`
}

func LoadConfigFromIni(fileName string) (Config, error) {
	CONFIG := new(Config)
	err := ini.MapTo(CONFIG, fileName)
	if err != nil {
		return *CONFIG, nil
	}
	return *CONFIG, err
}
