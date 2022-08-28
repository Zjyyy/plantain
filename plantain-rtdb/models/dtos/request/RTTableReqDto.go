package request

type RTTableReqDto struct {
	PID          string `json:"pid"`
	Value        string `json:"value"`
	ValueType    string `json:"valueType"`
	Des          string `json:"des"`
	Address      string `json:"address"`
	LimitUp      string `json:"limitUp"`
	LimitDown    string `json:"limitDown"`
	Level        uint   `json:"level"`
	AlarmDes     string `json:"alarmDes"`
	IsHistorical bool   `json:"isHistorical"`
}
