package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"qyyh-go/database"
	"qyyh-go/database/table"
	"regexp"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	database.MysqlConnInit()
	var list []table.Book
	err := database.DB().Debug().Find(&list).Error
	fmt.Println(err)
	for _, book := range list {
		book.Score = float32(book.Xiancao*5+book.Liangcao*4+book.Gancao*3+book.Kucao*2+book.Ducao*1) / float32(book.Xiancao+book.Liangcao+book.Gancao+book.Kucao+book.Ducao)
		database.DB().Debug().Updates(&book)
	}
	fmt.Println("done")
}

func Test2(t *testing.T) {
	s := "[CQ:image,file=aed5ed962d915dae5bf1fa9b28e67fe3.image,subType=11,url=https://gchat.qpic.cn/gchatpic_new/331767027/769964102-2493108392-AED5ED962D915DAE5BF1FA9B28E67FE3/0?term=2&amp;is_origin=0]"
	maps := subCQCode(s)
	base := getBase64(maps[0]["url"])
	fmt.Println(base)

}

func getBase64(url string) string {

	res, _ := http.Get(url)
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	return base64.StdEncoding.EncodeToString(data)
}

func subCQCode(cq string) []map[string]string {
	reg := regexp.MustCompile(`\[([^]\[\r\n]*)]`)
	var list []map[string]string

	for _, s := range reg.FindAllString(cq, -1) {
		ss := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", ""), ",")
		tmpMap := map[string]string{}
		for _, item := range ss {
			c := strings.Split(item, "=")
			if len(c) == 1 {
				tmpMap["type"] = c[0]
			} else {
				tmpMap[c[0]] = c[1]
			}
		}
		list = append(list, tmpMap)
	}
	return list
}

type Config struct {
	//本体数据库
	Mysqlhost   string //mysql地址
	Mysqluser   string //mysql用户名
	Mysqlpwd    string //mysql密码
	Mysqldbname string //mysql数据库名
}
