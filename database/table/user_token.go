package table

import (
	"github.com/google/uuid"
	"qyyh-go/database"
	"time"
)

type Token struct {
	Id        int       `json:"id" gorm:"column:id"`
	UserId    int       `json:"userid" gorm:"column:userid"`
	Token     string    `json:"token" gorm:"column:token"`
	LoginTime time.Time `json:"logintime" gorm:"column:logintime"`
}

func (m *Token) TableName() string {
	return "user_token"
}

func GetUserTokenByUserid(userid int) (data Token) {
	database.DB().Where("userid = ?", userid).Find(&data)
	return
}

func CreateUserToken(userid int) (data string, err error) {
	token := GetUserTokenByUserid(userid)
	data = uuid.NewString()
	if token.Id != 0 {
		err = database.DB().Where("id = ?", token.Id).Updates(&Token{
			UserId:    userid,
			Token:     data,
			LoginTime: time.Now(),
		}).Error
	} else {
		err = database.DB().Create(&Token{
			UserId:    userid,
			Token:     data,
			LoginTime: time.Now(),
		}).Error
	}
	return
}
