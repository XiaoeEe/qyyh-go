package config

import (
	"fmt"
	"os"
)

type Config struct {
	//本体数据库
	Mysqlhost string //mysql地址
	Mysqluser string //mysql用户名
	Mysqlpwd  string //mysql密码
	Mysqldb   string //mysql数据库名
}

func GetConfig() Config {
	goEnv := os.Getenv("GO_ENV")
	if goEnv != "local" && goEnv != "production" && goEnv != "ONLINE" {
		goEnv = "local"
	}

	localConfig := Config{
		Mysqlhost: "192.168.2.15:3306",
		Mysqluser: "root",
		Mysqlpwd:  "Zxc940429+++",
		Mysqldb:   "qyyh",
	}
	productionConfig := Config{
		Mysqlhost: "host:3306",
		Mysqluser: "root",
		Mysqlpwd:  "Zxc940429+++",
		Mysqldb:   "qyyh",
	}
	fmt.Printf("run in %s\n", goEnv)
	switch goEnv {
	case "local":
		return localConfig
	case "production":
		return productionConfig
	case "ONLINE":
		return Config{
			Mysqlhost: "qyyh.net:3306",
			Mysqluser: "root",
			Mysqlpwd:  "Zxc940429+++",
			Mysqldb:   "qyyh",
		}
	}
	return localConfig
}
