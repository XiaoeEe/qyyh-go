package table

import (
	"qyyh-go/database"
	"time"
)

type Roll struct {
	Id   int    `json:"id" gorm:"column:id"`
	QQ   int64  `json:"qq" gorm:"column:qq"`
	Date string `json:"date" gorm:"column:date"`
	Text string `json:"text" gorm:"column:text"`
}

func (m *Roll) TableName() string {
	return "robot_roll"
}

func GetRollByQQAndDate(qq int64, date string) (data Roll) {
	database.DB().Where("qq = ? && date = ?", qq, date).Find(&data)
	return
}

func CreateRoll(data Roll) error {
	return database.DB().Create(&data).Error
}

func CleanRoll() {
	database.DB().Table("robot_roll").Where("date != ?", time.Now().Format("20060102")).Delete(&Roll{})
}
