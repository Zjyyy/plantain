package main

import (
	"log"
	"plantain-common/common"
	"time"
)

//go build -o opcdriver.so -buildmode=plugin driver.go
type driver struct {
	rtdb common.IRTDB
	conf common.DriverConfigure
}

func (d *driver) Initialize(conf common.DriverConfigure, rtdb common.IRTDB) error {
	log.Println(">>>>>Initialize")
	d.conf = conf
	d.rtdb = rtdb
	return nil
}

func (d *driver) Do() error {
	log.Println(">>>>>>Do")
	d.rtdb.Write("Tag01", "1")
	d.rtdb.Read("Tag02")
	time.Sleep(time.Duration(2) * time.Second)
	return nil
}

var Driver driver
