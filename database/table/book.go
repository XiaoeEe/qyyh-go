package table

type Book struct {
	ID       int64   `json:"id" gorm:"column:id"`
	Title    string  `json:"title" gorm:"column:title"`
	Author   string  `json:"author" gorm:"column:author"`
	Xiancao  int     `json:"xiancao" gorm:"column:xiancao"`
	Liangcao int     `json:"liangcao" gorm:"column:liangcao"`
	Gancao   int     `json:"gancao" gorm:"column:gancao"`
	Kucao    int     `json:"kucao" gorm:"column:kucao"`
	Ducao    int     `json:"ducao" gorm:"column:ducao"`
	Score    float32 `json:"score" gorm:"column:score"`
}

func (m *Book) TableName() string {
	return "book"
}
