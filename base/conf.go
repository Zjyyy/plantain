package base

import (
	"github.com/go-ini/ini"
)

type System struct {
	Database string `ini:"database"`
}
type Config struct {
	System System `ini:"system"`
}

func LoadConfigFromIni(fileName string) (Config, error) {
	CONFIG := new(Config)
	err := ini.MapTo(CONFIG, fileName)
	if err != nil {
		return *CONFIG, nil
	}
	return *CONFIG, err
}
