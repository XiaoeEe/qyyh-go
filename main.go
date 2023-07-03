package main

import (
	"github.com/gin-gonic/gin"
	"qyyh-go/database"
	"qyyh-go/middleware"
	"qyyh-go/router"
	"qyyh-go/task"
	"time"
)

func main() {
	time.Local = time.FixedZone("CST", 8*3600) // 东八

	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	r.Use(middleware.Verify())
	r.Use(gin.Recovery())

	database.MysqlConnInit()
	router.Init(r)
	task.Init()
	if err := r.Run(":5702"); err != nil {
		panic(err)
	}
}
