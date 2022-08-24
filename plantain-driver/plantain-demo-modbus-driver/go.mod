module plantain/driver/modbus/demo

go 1.18
require plantain-common/common v0.0.0

replace plantain-common/common => ../../plantain-common
require (
	github.com/goburrow/serial v0.1.0 // indirect
	github.com/thinkgos/gomodbus/v2 v2.2.2 // indirect
)
