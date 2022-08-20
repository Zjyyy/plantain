package sqlite

import (
	"fmt"
	"log"
	"os"
	"plantain/base"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	SQLITE_DATABASE_NAME = "PlantinTestConfigurationDatabase.db"
	DRIVER_LIST_NAME     = "p_driver_list"
	DRIVER_TABLE_NAME    = "p_driver"
)

func GetConfigurationDatabaseHandler() *ConfigurationDatabaseHandler {
	db, dbErr := gorm.Open(sqlite.Open(SQLITE_DATABASE_NAME))
	if dbErr != nil {
		panic(fmt.Sprintf("打开SQLite连接错误：%v \n", dbErr))
	}

	return NewConfigurationDatabaseHandler(
		db, DRIVER_LIST_NAME, DRIVER_TABLE_NAME)
}

func CreateTestSqlite() {
	log.Printf("创建测试用Sqlite")
	_, err := os.Stat(SQLITE_DATABASE_NAME)

	db, dbErr := gorm.Open(sqlite.Open(SQLITE_DATABASE_NAME))
	if dbErr != nil {
		panic(fmt.Sprintf("打开SQLite连接错误：%v \n", dbErr))
	}

	handler := NewConfigurationDatabaseHandler(
		db, DRIVER_LIST_NAME, DRIVER_TABLE_NAME)

	if err != nil || os.IsNotExist(err) {
		log.Printf("当前项目下没有Plantain配置库，自动创建样例配置库\n")
		db.AutoMigrate(
			&base.RtTable{},
			&base.PDriverInDatabase{},
		)
		CreateMockData(handler)
	}
}

func DeleteTestSqlite() {
	log.Printf("移除测试用Sqlite")
	_, err := os.Stat(SQLITE_DATABASE_NAME)

	if err == nil || os.IsExist(err) {
		err = os.Remove(SQLITE_DATABASE_NAME)
		if err != nil {
			panic("删除测试用配置库失败")
		}
	}
}

func TestMain(m *testing.M) {
	log.Printf("开始ConfigurationDatabase测试")
	CreateTestSqlite()
	m.Run()
	DeleteTestSqlite()
	log.Printf("结束ConfigurationDatabase测试")
}

func TestLoadAllDriver(t *testing.T) {}

func TestReadAllRtTable(t *testing.T) {}

func TestRtTableCreateTable(t *testing.T) {
	handler := GetConfigurationDatabaseHandler()

	handler.CreateRTTable("rt_test_normal_create")
	if !handler.ExistRTTable("rt_test_normal_create") {
		t.Fatal("正常创建RTTable测试失败")
	}

	handler.CreateRTTable("")
	if !handler.ExistRTTable("") {
		t.Fatal("空值创建RTTable测试失败")
	}
}

func TestRtTableReadAddDropMethods(t *testing.T) {
	handler := GetConfigurationDatabaseHandler()
	arr, err := handler.ReadAllRtTable("noExist")
	if err == nil || len(arr) != 0 {
		t.Fatal("测试ReadAllRtTable读取不存在表失败 ")
	}

	handler.CreateRTTable("rt_test_read")
	if !handler.ExistRTTable("rt_test_read") {
		t.Fatal("CreateRTTable或ExistRTTable方法错误")
	}

	handler.AddRTTableItem("rt_test_read", &base.RtTable{
		PID:          "Tag01",
		Des:          "",
		Value:        "1",
		ValueType:    "int",
		Address:      "10001",
		LimitUp:      "10",
		LimitDown:    "-1",
		AlarmDes:     "",
		Level:        1,
		IsHistorical: true,
	})
	rtItem, err := handler.ReadAllRtTable("rt_test_read")
	if err != nil || len(rtItem) != 1 || rtItem[0].PID != "Tag01" {
		t.Fatal("AddRTTable或ReadAllRtTable方法错误")
	}
}

func TestDelRtTableItem(t *testing.T) {}
func TestDropRtTable(t *testing.T)    {}

func TestReadDriverList(t *testing.T)    {}
func TestAddDriverListItem(t *testing.T) {}
func TestDelDriverListItem(t *testing.T) {}
