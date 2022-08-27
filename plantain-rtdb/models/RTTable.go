package models

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
