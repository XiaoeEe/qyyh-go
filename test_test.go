package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"qyyh-go/database"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"strings"
	"sync"
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

func Test1(t *testing.T) {
	database.MysqlConnInit()
	rooms := table.GetLiveroomList()
	wg := sync.WaitGroup{}
	live := ""
	unlive := ""
	errored := ""
	s := ""
	for _, room := range rooms {
		wg.Add(1)
		go func(room table.Liveroom) {
			defer wg.Done()
			data := getBilibiliData(room.Uid)
			if data.Data.LiveRoom.RoomStatus == 0 {
				errored += fmt.Sprintf("%s的直播间:\n直播间状态获取错误\n", room.Nickname)
				return
			}
			url := strings.Split(data.Data.LiveRoom.Url, "?")[0]
			if data.Data.LiveRoom.LiveStatus == 0 {
				unlive += fmt.Sprintf("%s的直播间:\n%s: %s\n", room.Nickname, data.Data.LiveRoom.Title, url)
			} else {
				live += fmt.Sprintf("%s的直播间:\n%s: %s\n", room.Nickname, data.Data.LiveRoom.Title, url)
			}
		}(room)
	}
	wg.Wait()
	if live != "" {
		s += fmt.Sprintf("正在直播的直播间:\n%s\n\n", live)
	}
	if unlive != "" {
		s += fmt.Sprintf("没在直播的直播间:\n%s\n\n", unlive)
	}
	if errored != "" {
		s += fmt.Sprintf("获取失败的直播间:\n%s", errored)
	}
}

func getBilibiliData(uid string) model.BilibiliData {
	for i := 0; i < 100; i++ {
		data := model.BilibiliData{}
		client := http.Client{}
		get, err := http.NewRequest("GET", "https://api.bilibili.com/x/space/acc/info?mid="+uid, nil)
		if err != nil {
			return model.BilibiliData{}
		}
		get.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		get.Header.Set("Cache-Control", "no-cache")
		get.Header.Set("Connection", "keep-alive")
		get.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.62")
		do, err := client.Do(get)
		if err != nil {
			continue
		}
		readAll, err := io.ReadAll(do.Body)
		err = json.Unmarshal(readAll, &data)
		if err != nil {
			continue
		}
		return data
	}
	return model.BilibiliData{}
}
