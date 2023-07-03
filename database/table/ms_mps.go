package table

import (
	"fmt"
	"qyyh-go/database"
	"time"
)

type MPS struct {
	ID       int64  `json:"ID" gorm:"column:id"`
	Name     string `json:"Name" gorm:"column:name"`
	Level    int    `json:"Level" gorm:"column:level"`
	Points   int    `json:"Points" gorm:"column:points"`
	Culvert  int    `json:"Culvert" gorm:"column:culvert"`
	FlagRace int    `json:"FlagRace" gorm:"column:flagrace"`
	Date     string `json:"Date" gorm:"column:date"`
}

func (m *MPS) TableName() string {
	return "ms_mps"
}

func GetMPSDateList() (data []string) {
	database.DB().Table("ms_mps").Select("date").Group("date").Order("date desc").Scan(&data)
	return
}

func GetMPSByDate(date string) (data []MPS) {
	database.DB().Debug().Where("date = ?", date).Find(&data)
	return
}

func CreateMPS(data []MPS) error {
	return database.DB().CreateInBatches(&data, 17).Error
}

func DelMPS(ids []int64) error {
	return database.DB().Where("id in ?", ids).Delete(&MPS{}).Error
}

func GetMPSByNameAndDate(name, date string) (data MPS) {
	database.DB().Where("name = ? and date = ?", name, date).Find(&data)
	return
}

func GetFlagRaceByQQ(qq int64) int {
	sum := 0
	year, week := time.Now().ISOWeek()
	database.DB().Table("ms_mps").Select("SUM(flagrace)").Where("name in (select name FROM ms_player WHERE qq = ?) and date = ?", qq, fmt.Sprintf("%d-%d周", year, week-1)).Scan(&sum)
	return sum
}
