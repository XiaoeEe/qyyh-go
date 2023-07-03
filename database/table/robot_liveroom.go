package table

import "qyyh-go/database"

type Liveroom struct {
	Uid      string `json:"uid" gorm:"column:uid"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
	Flag     string `json:"flag" gorm:"column:flag"`
}

func (m *Liveroom) TableName() string {
	return "robot_liveroom"
}

func GetLiveroomList() (data []Liveroom) {
	database.DB().Find(&data)
	return
}

func SaveLiveroom(liveroom Liveroom) error {
	data := Liveroom{}
	database.DB().Where("uid = ?", liveroom.Uid).Find(&data)
	if data.Uid != "" {
		return database.DB().Updates(&liveroom).Error
	} else {
		return database.DB().Create(&liveroom).Error
	}
}

func DeleteLiveroom(uid string) error {
	return database.DB().Where("uid = ?", uid).Delete(&Liveroom{}).Error
}
