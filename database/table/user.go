package table

import (
	"qyyh-go/database"
)

type User struct {
	Id       int    `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	QQ       int64  `json:"qq" gorm:"column:qq"`
	Unionid  string `json:"unionid" gorm:"column:unionid"`
}

func (m *User) TableName() string {
	return "user"
}

func GetUserByUsernameAndPassword(username, password string) (data User) {
	database.DB().Where("username = ? and password = ?", username, password).Find(&data)
	return
}

func HasUserByUsername(username string) bool {
	var count int64
	database.DB().Table("user").Where("username = ?", username).Count(&count)
	return count > 0
}

func CreatUser(user User) (data User, err error) {
	err = database.DB().Create(&user).Error
	return
}

func GetUserByUnionid(unionid string) (data User) {
	database.DB().Where("unionid = ?", unionid).Find(&data)
	return
}

func GetUserByToken(token string) (data User) {
	database.DB().Where("id = (select userid from user_token where token = ?)", token).Find(&data)
	return
}

func UpdateUser(user User) error {
	return database.DB().Updates(&user).Error
}

func GetUserByQQ(qq int64) (data User) {
	database.DB().Where("qq = ?", qq).Find(&data)
	return
}
