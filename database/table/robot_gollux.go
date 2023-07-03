package table

import "qyyh-go/database"

type Gollux struct {
	QQ int64 `json:"qq" gorm:"column:qq"`
}

func (m *Gollux) TableName() string {
	return "robot_gollux"
}

func CreateGollux(data Gollux) error {
	return database.DB().Create(&data).Error
}

func DelGollux(qq int64) error {
	return database.DB().Where("qq = ?", qq).Delete(&Gollux{}).Error
}

func GetGolluxList() (data []Gollux) {
	database.DB().Find(&data)
	return
}
