package table

import "qyyh-go/database"

type Location struct {
	Location string `json:"location" gorm:"column:location"`
}

func (m *Location) TableName() string {
	return "robot_location"
}

func GetLocations() (data []string) {
	database.DB().Table("robot_location").Select("location").Scan(&data)
	return
}
