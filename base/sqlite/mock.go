package sqlite

import (
	"plantain/base"

	"gorm.io/gorm"
)

func CreateMockData(db *gorm.DB) {
	CreateMockDriverList(db)
	CreateMockRTTable(db)
}

func CreateMockDriverList(db *gorm.DB) {
	AddDriverListItem(db, &base.PDriverInDatabase{
		Id:            1,
		DriverName:    "Demo1",
		Version:       "0.0.1",
		DriverDllPath: "",
		ConnStr:       "dev-asyncua;4840",
		Setting:       "setting1;setting2",
		Des:           "opc ua demo",
		RtTableName:   "rt_opcdemo",
	})
	AddDriverListItem(db, &base.PDriverInDatabase{
		Id:            2,
		DriverName:    "Demo2",
		Version:       "0.0.1",
		DriverDllPath: "",
		ConnStr:       "dev-asyncua;4840",
		Setting:       "setting1;setting2",
		Des:           "modbus demo",
		RtTableName:   "rt_modbusdemo",
	})
}

func CreateMockRTTable(db *gorm.DB) {
	CreateMockModbusRTTable(db)
	CreateMockOPCRTTable(db)
}

func CreateMockModbusRTTable(db *gorm.DB) {
	CreateRTTable(db, "rt_modbusdemo")
	AddRTTableItem(db, "rt_modbusdemo", &base.RtTable{
		PID:          "Tag01",
		Value:        "1",
		ValueType:    "int",
		Address:      "10001",
		LimitUp:      "10",
		LimitDown:    "-1",
		Level:        1,
		IsHistorical: true,
	})
	AddRTTableItem(db, "rt_modbusdemo", &base.RtTable{
		PID:          "Tag02",
		Value:        "0",
		ValueType:    "float",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-10",
		Level:        1,
		IsHistorical: false,
	})
	AddRTTableItem(db, "rt_modbusdemo", &base.RtTable{
		PID:          "Tag03",
		Value:        "false",
		ValueType:    "bool",
		Address:      "10003",
		LimitUp:      "true",
		LimitDown:    "false",
		Level:        1,
		IsHistorical: true,
	})
}

func CreateMockOPCRTTable(db *gorm.DB) {
	CreateRTTable(db, "rt_opcdemo")
	AddRTTableItem(db, "rt_opcdemo", &base.RtTable{
		PID:          "Tag01",
		Value:        "1",
		ValueType:    "int",
		Address:      "10001",
		LimitUp:      "10",
		LimitDown:    "-1",
		Level:        1,
		IsHistorical: true,
	})
	AddRTTableItem(db, "rt_opcdemo", &base.RtTable{
		PID:          "Tag02",
		Value:        "0",
		ValueType:    "float",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-10",
		Level:        1,
		IsHistorical: false,
	})
	AddRTTableItem(db, "rt_opcdemo", &base.RtTable{
		PID:          "Tag03",
		Value:        "false",
		ValueType:    "bool",
		Address:      "10003",
		LimitUp:      "true",
		LimitDown:    "false",
		Level:        1,
		IsHistorical: true,
	})
}
