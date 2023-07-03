package table

type Daily struct {
	Id         int     `json:"id" gorm:"column:id"`
	MSName     string  `json:"msname" gorm:"column:msname"`
	QQ         int64   `json:"qq" gorm:"column:qq"`
	Group      int64   `json:"group" gorm:"column:groupid"`
	Exp        int64   `json:"exp" gorm:"column:exp"`
	EXPPercent float64 `json:"exppercent" gorm:"column:exppercent"`
	Level      int64   `json:"level" gorm:"column:level"`
	Date       string  `json:"date" gorm:"column:date"`
}

func (m *Daily) TableName() string {
	return "robot_daily"
}
