package core

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

func TestBuildMemoryStructure(t *testing.T) {
	pDriver := createMockPDriver()

	cacheArr := BuildMemoryStructure(pDriver)

	//Set Get为go-cache提供的方法
	cache1 := cacheArr["Name1"]
	cache1.Set("Tag01", 1, 0)
	result1, _ := cache1.Get("Tag01")

	cache2 := cacheArr["Name2"]
	cache2.Set("Tag01", "2", 0)
	result2, _ := cache2.Get("Tag02")

	if result1 != 1 || result2 != "2" {
		t.Fatal("没有正确创建互不影响的内存片")
	}
}
