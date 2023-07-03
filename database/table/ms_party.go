package table

import (
	"qyyh-go/database"
	"time"
)

type Party struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Leader   string `json:"leader" gorm:"column:leader"`
	Member   string `json:"member" gorm:"column:member"`
	Boss     string `json:"boss" gorm:"column:boss"`
	Count    int    `json:"count" gorm:"column:count"`
	Level    int64  `json:"level" gorm:"column:level"`
	Point    int    `json:"point" gorm:"column:point"`
	Remarks  string `json:"remarks" gorm:"column:remarks"`
	Type     int    `json:"type" gorm:"column:type"`
	Time     string `json:"time" gorm:"column:time"`
	Channel  int    `json:"channel" gorm:"column:channel"`
	CreateBy int    `json:"createBy" gorm:"column:createby"`
}

func (m *Party) TableName() string {
	return "ms_party"
}

func CreateParty(data Party) error {
	return database.DB().Create(&data).Error
}

func GetPartyListByType(t int) (data []Party) {
	if t == 1 {
		t1, _ := time.ParseDuration("-120m")
		database.DB().Where("type = ? and time > ?", t, time.Now().Add(t1).Format("2006-01-02 15:04:05")).Order("time").Find(&data)
	} else {
		database.DB().Where("type = ?", t).Find(&data)
	}
	return
}

func DelParty(id int64) error {
	return database.DB().Where("id = ?", id).Delete(&Party{}).Error
}

func GetPartyById(id int64) (data Party) {
	database.DB().Where("id = ?", id).Find(&data)
	return
}

func UpdateParty(data Party) error {
	return database.DB().Save(&data).Error
}

func GetUsePointByQQ(qq int64) (point int) {
	var partyList []Party
	database.DB().Where("YEARWEEK(time) = YEARWEEK(NOW())").Find(&partyList)
	players := GetMSPlayerListByQQ(qq)

	for _, party := range partyList {
		for _, player := range players {
			var find int
			database.DB().Raw("SELECT FIND_IN_SET(?, ?)", player.Name, party.Member).Scan(&find)
			if find != 0 {
				point += party.Point
				break
			}
		}
	}

	return
}

func GetPartyCountByQQ(qq int64) (count int64) {
	database.DB().Table("ms_party").Where("leader in (select name from ms_player where qq = ?) and member != ?", qq, "").Count(&count)
	return
}
