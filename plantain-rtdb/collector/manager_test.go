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
		Address:      "10001",
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
	//后续考虑支持插拔式使用loadPlugin方法时补上
}

var CreateCommonDriverConfigure = createCommonDriverConfigure

func TestCreateCommonDriverConfigure(t *testing.T) {
	pDriver := createMockPDriver()
	configure := createCommonDriverConfigure(pDriver[0])

	if configure.Id != 1 ||
		configure.ConnStr != "127.0.0.1:3000" ||
		configure.DriverName != "Name1" ||
		configure.RtPoint["Tag01"] != "10001" {
		t.Fatal("创建驱动配置的数据结构异常")
	}
}
