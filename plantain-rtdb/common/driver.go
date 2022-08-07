package common

type DriverConfigure struct {
	Id         uint
	DriverName string
	ConnStr    string
	RtPoint    map[string]string
}

type IDriver interface {
	Do() error
	Initialize(conf DriverConfigure, rtdb IRTDB) error
}
