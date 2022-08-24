package main

import (
	"log"
	"plantain-common/common"
	"time"

	modbus "github.com/thinkgos/gomodbus/v2"
)

type driver struct {
	rtdb         common.IRTDB
	conf         common.DriverConfigure
	modbusClient modbus.Client
}

func (d *driver) Initialize(conf common.DriverConfigure, rtdb common.IRTDB) error {
	log.Println(">>>>>Initialize")
	d.conf = conf
	d.rtdb = rtdb
	p := modbus.NewTCPClientProvider(d.conf.ConnStr)
	d.modbusClient = modbus.NewClient(p)

	return nil
}

func (d *driver) Do() error {
	if !d.modbusClient.IsConnected() {
		err := d.modbusClient.Connect()
		if err != nil {
			return err
		}
	}
	results, err := d.modbusClient.ReadHoldingRegisters(1, 0, 1)
	if err != nil {
		return err
	}
	d.rtdb.Write("Tag01", results[0])

	val := d.rtdb.Read("Tag01")
	log.Printf(">>Tag01:%v", val)

	time.Sleep(time.Duration(2) * time.Second)
	return nil
}

var Driver driver
