package collector

import (
	"plantain/base"
	"testing"
)

func createMockPDriver() []base.PDriver {
	pDriver := make([]base.PDriver, 2)
	rtTable := make([]base.RtTable, 2)
	rtTable[0] = base.RtTable{
		PID:          "Tag01",
		Value:        "1",
		ValueType:    "int",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-1",
		Level:        1,
		IsHistorical: true,
	}
	rtTable[1] = base.RtTable{
		PID:          "Tag02",
		Value:        "2",
		ValueType:    "int",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-1",
		Level:        1,
		IsHistorical: true,
	}
	pDriver[0] = base.PDriver{
		Id:            1,
		DriverName:    "Name1",
		Version:       "0.0.1",
		DriverDllPath: "/home/test",
		ConnStr:       "127.0.0.1:3000",
		Setting:       "timeout=1",
		Des:           "des1",
		RtTable:       rtTable,
	}
	pDriver[1] = base.PDriver{
		Id:            2,
		DriverName:    "Name2",
		Version:       "0.0.1",
		DriverDllPath: "/home/test",
		ConnStr:       "127.0.0.1:8000",
		Setting:       "timeout=1",
		Des:           "des2",
		RtTable:       rtTable,
	}
	return pDriver
}

func TestInitCollector(t *testing.T) {
}

var LoadDriverPlugin = loadDriverPlugin

func TestLoadDriverPlugin(t *testing.T) {

}

var CreateCommonDriverConfigure = createCommonDriverConfigure

func TestCreateCommonDriverConfigure(t *testing.T) {

}
