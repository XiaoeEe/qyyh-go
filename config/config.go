package config

import (
	"os"
)

type Config struct {
	//本体数据库
	Mysqlhost   string //mysql地址
	Mysqluser   string //mysql用户名
	Mysqlpwd    string //mysql密码
	Mysqldbname string //mysql数据库名
}

func GetConfig() Config {

	config := Config{
		Mysqlhost:   os.Getenv("Mysqlhost"),
		Mysqluser:   os.Getenv("Mysqluser"),
		Mysqlpwd:    os.Getenv("Mysqlpwd"),
		Mysqldbname: os.Getenv("Mysqldbname"),
	}

	return config
}
