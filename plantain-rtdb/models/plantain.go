package models

type AlarmMessage struct {
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
