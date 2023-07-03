package table

import (
	"qyyh-go/database"
	"time"
)

type MSPlayer struct {
	Name     string    `json:"name" gorm:"column:name"`
	Class    string    `json:"class" gorm:"column:class"`
	Level    int64     `json:"level" gorm:"column:level"`
	Img      string    `json:"img" gorm:"column:img"`
	QQ       int64     `json:"qq" gorm:"column:qq"`
	IsMain   bool      `json:"ismain" gorm:"column:ismain"`
	Datatime time.Time `json:"datatime" gorm:"column:datatime"`
}

func (m *MSPlayer) TableName() string {
	return "ms_player"
}

func (m *MSPlayer) Equals(e MSPlayer) bool {
	return m.Name == e.Name && m.Class == e.Class && m.Level == e.Level && m.Img == e.Img && m.Datatime == e.Datatime && m.QQ == e.QQ && m.IsMain == e.IsMain
}

func GetMSPlayer(name string) (data MSPlayer) {
	database.DB().Where("name = ?", name).Find(&data)
	return
}

func GetMSPlayerList() (data []MSPlayer) {
	database.DB().Find(&data)
	return
}

func GetMSPlayerListByQQ(qq int64) (data []MSPlayer) {
	database.DB().Where("qq = ?", qq).Find(&data)
	return
}

func SaveMSPlayer(player MSPlayer) error {
	p := GetMSPlayer(player.Name)
	if p.Equals(player) {
		return nil
	}
	if p.Name != "" {
		return database.DB().Debug().Where("name = ?", player.Name).Updates(&player).Error
	} else {
		return database.DB().Debug().Create(&player).Error
	}
}

func DeleteMSPlayer(player MSPlayer) error {
	return database.DB().Where("name = ?", player.Name).Delete(&player).Error
}

func GetMainPlayerByQQ(qq int64) (data MSPlayer) {
	database.DB().Where("qq = ? and ismain = 1", qq).Find(&data)
	return
}

func CleanPlayer(qqs []int64) {
	database.DB().Where("qq not in (?)", qqs).Delete(&MSPlayer{})
}
