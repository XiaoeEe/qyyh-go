package table

import "qyyh-go/database"

type Entry struct {
	Name  string `json:"name" gorm:"column:name"`
	Value string `json:"value" gorm:"column:value"`
}

func (m *Entry) TableName() string {
	return "robot_entry"
}

func GetEntry(name string) (data Entry) {
	database.DB().Where("name = ?", name).Find(&data)
	return
}

func GetEntryList() (data []Entry) {
	database.DB().Find(&data)
	return
}

func SaveEntry(entry Entry) error {
	data := GetEntry(entry.Name)
	if data.Name == "" {
		return database.DB().Create(&entry).Error
	} else {
		return database.DB().Debug().Where("name = ?", entry.Name).Updates(&entry).Error
	}
}

func GetEntryListByLikeName(name string) (data []Entry) {
	database.DB().Where("name like ?", "%"+name+"%").Find(&data)
	return
}

func DeleteEntry(name string) error {
	return database.DB().Where("name = ?", name).Delete(&Entry{}).Error
}
