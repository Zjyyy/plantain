package models

func CreateMockData() {
	AddCollectorItemInList(&Collector{
		Id:            1,
		CollectorName: "MockModbus1",
		Version:       "0.0.1",
		DllPath:       "/home/dev/plantain-main/plantain-driver/plantain-demo-modbus-driver/modbusdriver.so",
		ConnStr:       "mock-modbus-1:123",
		Setting:       "setting1;setting2",
		Des:           "modbus demo",
		RtTableName:   "rt_modbusdemo",
	})
}
