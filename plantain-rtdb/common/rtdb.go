package common

type IRTDB interface {
	Write(pid string, value string) bool
	Read(pid string) string
}
