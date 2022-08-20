package sqlite

import (
	"plantain/base"
)

func CreateMockData(handler *ConfigurationDatabaseHandler) {
	CreateMockDriverList(handler)
	CreateMockRTTable(handler)
}

func CreateMockDriverList(handler *ConfigurationDatabaseHandler) {
	handler.AddDriverListItem(&base.PDriverInDatabase{
		Id:            1,
		DriverName:    "Demo1",
		Version:       "0.0.1",
		DriverDllPath: "/home/dev/plantain-driver/plantain-demo-opc-driver/opcdriver.so",
		ConnStr:       "dev-asyncua;4840",
		Setting:       "setting1;setting2",
		Des:           "opc ua demo",
		RtTableName:   "rt_opcdemo",
	})
	// AddDriverListItem(db, &base.PDriverInDatabase{
	// 	Id:            2,
	// 	DriverName:    "Demo2",
	// 	Version:       "0.0.1",
	// 	DriverDllPath: "",
	// 	ConnStr:       "dev-asyncua;4840",
	// 	Setting:       "setting1;setting2",
	// 	Des:           "modbus demo",
	// 	RtTableName:   "rt_modbusdemo",
	// })
}

func CreateMockRTTable(handler *ConfigurationDatabaseHandler) {
	CreateMockModbusRTTable(handler)
	CreateMockOPCRTTable(handler)
}

func CreateMockModbusRTTable(handler *ConfigurationDatabaseHandler) {
	handler.CreateRTTable("rt_modbusdemo")
	handler.AddRTTableItem("rt_modbusdemo", &base.RtTable{
		PID:          "Tag01",
		Des:          "",
		Value:        "1",
		ValueType:    "int",
		Address:      "10001",
		LimitUp:      "10",
		LimitDown:    "-1",
		AlarmDes:     "",
		Level:        1,
		IsHistorical: true,
	})
	handler.AddRTTableItem("rt_modbusdemo", &base.RtTable{
		PID:          "Tag02",
		Des:          "",
		Value:        "1.3",
		ValueType:    "float",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-10",
		AlarmDes:     "",
		Level:        1,
		IsHistorical: false,
	})
	handler.AddRTTableItem("rt_modbusdemo", &base.RtTable{
		PID:          "Tag03",
		Des:          "",
		Value:        "false",
		ValueType:    "bool",
		Address:      "10003",
		LimitUp:      "true",
		LimitDown:    "false",
		Level:        1,
		AlarmDes:     "",
		IsHistorical: true,
	})
}

func CreateMockOPCRTTable(handler *ConfigurationDatabaseHandler) {
	handler.CreateRTTable("rt_opcdemo")
	handler.AddRTTableItem("rt_opcdemo", &base.RtTable{
		PID:          "Tag01",
		Des:          "",
		Value:        "1",
		ValueType:    "int",
		Address:      "10001",
		LimitUp:      "10",
		LimitDown:    "-1",
		Level:        1,
		AlarmDes:     "",
		IsHistorical: true,
	})
	handler.AddRTTableItem("rt_opcdemo", &base.RtTable{
		PID:          "Tag02",
		Des:          "",
		Value:        "0",
		ValueType:    "float",
		Address:      "10002",
		LimitUp:      "10",
		LimitDown:    "-10",
		Level:        1,
		AlarmDes:     "",
		IsHistorical: false,
	})
	handler.AddRTTableItem("rt_opcdemo", &base.RtTable{
		PID:          "Tag03",
		Des:          "",
		Value:        "false",
		ValueType:    "bool",
		Address:      "10003",
		LimitUp:      "true",
		LimitDown:    "false",
		Level:        1,
		AlarmDes:     "",
		IsHistorical: true,
	})
}
