package robotService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"strings"
	"sync"
)

func getLive(message model.CQMessage) {
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
	if s != "" {
		SendGroupMsg(message, s)
	} else {
		SendGroupMsg(message, "直播间为空")
	}
}

func addLiveroom(message model.CQMessage, nickname, uid string) {
	data := getBilibiliData(uid)
	if data.Code != 0 || data.Data.LiveRoom.RoomStatus == 0 {
		SendGroupMsg(message, "目前仅支持bilibili直播间登记, 请检查uid是否正确")
		return
	}
	if err := table.SaveLiveroom(table.Liveroom{
		Uid:      uid,
		Nickname: nickname,
		Flag:     "B站",
	}); err != nil {
		SendGroupMsg(message, err.Error())
	} else {
		SendGroupMsg(message, "直播间设置完成")
	}
}

func delLiveRoom(message model.CQMessage, uid string) {
	if err := table.DeleteLiveroom(uid); err != nil {
		SendGroupMsg(message, err.Error())
	} else {
		SendGroupMsg(message, "删除完成")
	}
}

func getBilibiliData(uid string) model.BilibiliData {
	for i := 0; i < 5; i++ {
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
