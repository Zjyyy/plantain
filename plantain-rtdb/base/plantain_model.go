package base

type PDriver struct {
	Id            uint
	DriverName    string
	Version       string
	DriverDllPath string
	ConnStr       string
	Setting       string
	Des           string
	RtTable       []RtTable
}

type PDriverInDatabase struct {
	Id            uint
	DriverName    string
	Version       string
	DriverDllPath string
	ConnStr       string
	Setting       string
	Des           string
	RtTableName   string
}

func (PDriverInDatabase) TableName() string {
	return "p_driver_lists"
}

type RtTable struct {
	PID          string
	Value        string
	ValueType    string
	Des          string
	Address      string
	LimitUp      string
	LimitDown    string
	Level        uint
	AlarmDes     string
	IsHistorical bool
}

type AlarmHistoryMessage struct {
	Table     string
	PID       string
	Des       string
	AlarmDes  string
	ValueType string
	Value     string
}

type HistoricalMessage struct {
	Table     string
	PID       string
	Value     string
	ValueType string
	Des       string
}
