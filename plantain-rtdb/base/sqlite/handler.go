package sqlite

import (
	"errors"
	"log"
	"plantain/base"

	"gorm.io/gorm"
)

type ConfigurationDatabaseHandler struct{ *configurationDatabase }
type configurationDatabase struct {
	db                  *gorm.DB
	driverListTableName string
	driverTableName     string
}

func NewConfigurationDatabaseHandler(db *gorm.DB, listTableName string, driverTableName string) *ConfigurationDatabaseHandler {
	return &ConfigurationDatabaseHandler{
		&configurationDatabase{
			db:                  db,
			driverListTableName: listTableName,
			driverTableName:     driverTableName,
		},
	}
}

func (self *configurationDatabase) LoadAllDriver() ([]base.PDriver, error) {
	dListFromDatabase, err := self.ReadDriverList()
	if err != nil {
		return nil, err
	}

	dList := make([]base.PDriver, len(dListFromDatabase))
	for index, item := range dListFromDatabase {
		rtList, err := self.ReadAllRtTable(item.RtTableName)
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

func (self *configurationDatabase) ReadAllRtTable(name string) ([]base.RtTable, error) {
	var rtList []base.RtTable
	result := self.db.Table(name).Find(&rtList)
	return rtList, result.Error
}

func (self *configurationDatabase) ExistRTTable(tableName string) bool {
	return self.db.Migrator().HasTable(tableName)
}

func (self *configurationDatabase) CreateRTTable(tableName string) error {
	//rt_tables是中间表，gorm不支持直接以表命名方式创建表
	if self.ExistRTTable("rt_tables") {
		log.Printf("must clean rt_tables.")
		self.db.Migrator().DropTable("rt_tables")
	}
	if hasTable := self.db.Migrator().HasTable(tableName); hasTable {
		return errors.New("driver table has exist.")
	}
	log.Printf("start create rt_tables as temp table.")

	if err := self.db.Migrator().CreateTable(&base.RtTable{}); err != nil {
		return err
	}
	log.Printf("start rename rt_tables to %s.", tableName)

	if err := self.db.Migrator().RenameTable("rt_tables", tableName); err != nil {
		self.db.Migrator().DropTable("rt_tables")
		return err
	}

	return nil
}

func (self *configurationDatabase) AddRTTableItem(tableName string, item *base.RtTable) error {
	return self.db.Table(tableName).Create(item).Error
}

func (self *configurationDatabase) DeleteRTTableItem(tableName string, pid string) error {
	return self.db.Table(tableName).Delete(&base.RtTable{}, "p_id LIKE ?", "%"+pid+"%").Error
}

func (self *configurationDatabase) DropRTTable(tableName string) error {
	if hasTable := self.db.Migrator().HasTable(tableName); hasTable {
		if err := self.db.Migrator().DropTable(tableName); err != nil {
			return err
		}
		log.Printf("Remove %v success.\n", tableName)
	} else {
		log.Printf("TableName: %v is not exist.\n", tableName)
	}
	return nil
}

func (self *configurationDatabase) ReadDriverList() ([]base.PDriverInDatabase, error) {
	var dList []base.PDriverInDatabase
	result := self.db.Table(self.driverListTableName).Find(&dList)
	return dList, result.Error
}
func (self *configurationDatabase) AddDriverListItem(item *base.PDriverInDatabase) error {
	return self.db.Create(item).Error
}
func (self *configurationDatabase) DelDriverListItem(name string) error {
	return self.db.Delete(&base.PDriverInDatabase{}, "driver_name LIKE ?", "%"+name+"%").Error
}
