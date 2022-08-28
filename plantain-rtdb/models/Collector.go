package models

type Collector struct {
	Id            uint   `json:"id"`
	CollectorName string `json:"collectorName" gorm:"unique"`
	Version       string `json:"version"`
	DllPath       string `json:"dllPath"`
	ConnStr       string `json:"connStr"`
	Setting       string `json:"setting"`
	Des           string `json:"des"`
	RtTableName   string `json:"rtTableName" gorm:"unique"`
}

const CollectorsTableName = "conf_collectors_list"

func (Collector) TableName() string {
	return CollectorsTableName
}

func GetAllCollectorList() ([]Collector, error) {
	var cList []Collector
	result := db.Table(CollectorsTableName).Find(&cList)
	return cList, result.Error
}

func GetCollectorByCollectorName(name string) (Collector, error) {
	var collector Collector
	result := db.Table(CollectorsTableName).
		Where("collector_name LIKE ?", "%"+name+"%").
		First(&collector)
	return collector, result.Error
}

func GetCollectorById(id int) (Collector, error) {
	var collector Collector
	result := db.Table(CollectorsTableName).
		Where("id = ?", id).
		First(&collector)
	return collector, result.Error
}

func AddCollectorItemInList(item *Collector) error {
	return db.Table(CollectorsTableName).Create(item).Error
}

func DelCollectorItemInListByName(collectorName string) error {
	return db.
		Where("collector_name LIKE ?", "%"+collectorName+"%").
		Delete(&Collector{}).Error
}

func DelCollectorItemInListById(id int) error {
	return db.
		Where("id = ?", id).
		Delete(&Collector{}).Error
}

func UpdateCollectorItemInListByName(collectorName string, data *Collector) error {
	maps := make(map[string]interface{})
	maps["collector_name"] = data.CollectorName
	maps["version"] = data.Version
	maps["dll_path"] = data.DllPath
	maps["conn_str"] = data.ConnStr
	maps["setting"] = data.ConnStr
	maps["des"] = data.Des
	maps["rt_table_name"] = data.RtTableName

	return db.Table(CollectorsTableName).
		Where("collector_name LIKE ?", "%"+collectorName+"%").
		Updates(&maps).Error
}
func UpdateCollectorItemInListById(id int, data *Collector) error {
	maps := make(map[string]interface{})
	maps["collector_name"] = data.CollectorName
	maps["version"] = data.Version
	maps["dll_path"] = data.DllPath
	maps["conn_str"] = data.ConnStr
	maps["setting"] = data.ConnStr
	maps["des"] = data.Des
	maps["rt_table_name"] = data.RtTableName
	return db.Table(CollectorsTableName).
		Where("id = ?", id).
		Updates(&maps).Error
}