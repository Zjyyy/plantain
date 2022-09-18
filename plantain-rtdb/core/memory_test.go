package core

import (
	"plantain/models"
	"testing"
)

func TestMemoryBlock(t *testing.T) {
	collector := models.Collector{
		Id:            1,
		CollectorName: "Test",
		Version:       "",
		DllPath:       "",
		ConnStr:       "",
		Setting:       "",
		Des:           "",
		RtTableName:   "test",
	}
	rtTable := models.RtTable{
		PID:          "t01",
		Value:        "",
		ValueType:    "",
		Des:          "",
		Address:      "",
		LimitUp:      "",
		LimitDown:    "",
		Level:        1,
		AlarmDes:     "",
		IsHistorical: false,
	}
	rtTableSet := []models.RtTable{rtTable}

	collectorArr := []models.CollectorWithRtTable{}
	collectorArr = append(collectorArr, models.CollectorWithRtTable{
		Collector:  collector,
		RtTableSet: rtTableSet,
	})
	memoryBlock := NewMemoryBlock()
	memoryBlock.BuildMemoryBlockSet(&collectorArr)
	if memoryBlock.CountMemoryBlockSet() != 1 {
		t.Fatal("CountMemoryBlockSet方法错误")
	}
	memoryBlock.WriteInCache("Test", "t01", "2")
	val, err := memoryBlock.ReadFromCache("Test", "t01")
	if err != nil {
		t.Fatal("ReadFromCache方法错误")
	}
	if val != "2" {
		t.Fatal("WritevalInCache方法错误")
	}
}

func TestGetCache(t *testing.T) {

}

func TestReadAndWrite(t *testing.T) {

}

func TestCountMemoryBlockSet(t *testing.T) {

}
