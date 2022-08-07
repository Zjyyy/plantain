package sqlite

import (
	"errors"
	"log"
	"plantain/base"

	"gorm.io/gorm"
)

const SQLiteName = "rtdbConfBase"
const PDriverListTableName = "p_driver_lists"
const PDriverTableName = "p_drivers"

func LoadAllDriver(db *gorm.DB) ([]base.PDriver, error) {
	dListFromDatabase, err := ReadDriverList(db)
	if err != nil {
		return nil, err
	}

	dList := make([]base.PDriver, len(dListFromDatabase))
	for index, item := range dListFromDatabase {
		rtList, err := ReadAllRtTable(db, item.RtTableName)
		if err != nil {
			return nil, err
		}
		var tempDriver base.PDriver
		tempDriver.Id = item.Id
		tempDriver.DriverName = item.DriverName
		tempDriver.Version = item.Version
		tempDriver.DriverDllPath = item.DriverDllPath
		tempDriver.ConnStr = item.ConnStr
		tempDriver.Setting = item.Setting
		tempDriver.Des = item.Des
		tempDriver.RtTable = rtList

		dList[index] = tempDriver
	}
	return dList, nil
}

func ReadAllRtTable(db *gorm.DB, name string) ([]base.RtTable, error) {
	var rtList []base.RtTable
	result := db.Table(name).Find(&rtList)
	return rtList, result.Error
}
func CreateRTTable(db *gorm.DB, tableName string) error {
	//rt_tables是中间表，gorm不支持直接以表命名方式创建表
	if hasTable := db.Migrator().HasTable("rt_tables"); hasTable {
		log.Printf("must clean rt_tables.")
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
func AddRTTableItem(db *gorm.DB, tableName string, item *base.RtTable) error {
	return db.Table(tableName).Create(item).Error
}
func DeleteRTTableItem(db *gorm.DB, tableName string, pid string) error {
	return db.Table(tableName).Delete(&base.RtTable{}, "p_id LIKE ?", "%"+pid+"%").Error
}
func DropRTTable(db *gorm.DB, tableName string) error {
	if hasTable := db.Migrator().HasTable(tableName); hasTable {
		if err := db.Migrator().DropTable(tableName); err != nil {
			return err
		}
		log.Printf("Remove %v success.\n", tableName)
	} else {
		log.Printf("TableName: %v is not exist.\n", tableName)
	}
	return nil
}

func ReadDriverList(db *gorm.DB) ([]base.PDriverInDatabase, error) {
	var dList []base.PDriverInDatabase
	result := db.Table(PDriverListTableName).Find(&dList)
	return dList, result.Error
}
func AddDriverListItem(db *gorm.DB, item *base.PDriverInDatabase) error {
	return db.Create(item).Error
}
func DelDriverListItem(db *gorm.DB, name string) error {
	return db.Delete(&base.PDriverInDatabase{}, "driver_name LIKE ?", "%"+name+"%").Error
}
