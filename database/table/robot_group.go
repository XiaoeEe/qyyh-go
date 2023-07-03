package table

import "qyyh-go/database"

type Group struct {
	Group      int64  `json:"group" gorm:"column:groupid"`
	Name       string `json:"name" gorm:"column:name"`
	WelcomeMsg string `json:"welcomemsg" gorm:"column:welcomemsg"`
}

func (m *Group) TableName() string {
	return "robot_group"
}

func GetGroupByGroupId(id int64) (data Group) {
	database.DB().Where("groupid = ?", id).Find(&data)
	return
}
