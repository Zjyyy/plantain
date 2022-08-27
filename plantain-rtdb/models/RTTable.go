package models

import (
	"errors"
	"fmt"
	"log"
	"plantain/base"
)

type RtTable struct {
	PID          string `json:"pid" gorm:"column:pid" gorm:"unique"`
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

func ExistRTTable(tableName string) bool {
	return db.Migrator().HasTable(tableName)
}

func GetRTTable(tableName string) ([]RtTable, error) {
	if !ExistRTTable(tableName) {
		return nil, errors.New("tablse is not exist")
	}

	var rtTable []RtTable
	result := db.Table(tableName).Find(&rtTable)

	return rtTable, result.Error
}

func CreateRTTable(tableName string) error {
	// rt_tables是中间表，gorm不支持直接以表命名方式创建表
	if ExistRTTable("rt_tables") {
		db.Migrator().DropTable("rt_tables")
	}
	if hasTable := db.Migrator().HasTable(tableName); hasTable {
		return errors.New("driver table has exist.")
	}
	log.Printf("start create rt_tables as temp table.")

	if err := db.Migrator().CreateTable(&base.RtTable{}); err != nil {
		return err
	}
	log.Printf("start rename rt_tables to %s.", tableName)

	if err := db.Migrator().RenameTable("rt_tables", tableName); err != nil {
		db.Migrator().DropTable("rt_tables")
		return err
	}

	return nil
}

func AddItemInRTTable(tableName string, item *RtTable) error {
	return db.Table(tableName).Create(item).Error
}

func DeleteItemInRTTableByPID(tableName string, pid string) error {
	return db.Table(tableName).
		Delete(
			&base.RtTable{},
			"pid LIKE ?",
			"%"+pid+"%").Error
}

func DropRTTable(tableName string) error {
	if hasTable := db.Migrator().HasTable(tableName); hasTable {
		if err := db.Migrator().DropTable(tableName); err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("TableName: %v is not exist.\n", tableName))
	}
	return nil
}
