package common

type IRTDB interface {
	Write(pid string, value interface{}) bool
	Read(pid string) interface{}
}
