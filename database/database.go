package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qyyh-go/config"
)

var db *gorm.DB

func MysqlConnInit() {
	c := config.GetConfig()

	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", c.Mysqluser, c.Mysqlpwd, c.Mysqlhost, c.Mysqldbname)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func DB() *gorm.DB {
	return db
}
