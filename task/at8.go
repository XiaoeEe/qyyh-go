package task

import (
	"fmt"
	"log"
	"os"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/service/robotSerice"
	"qyyh-go/utils"
	"sync"
	"time"
)

func At8() {
	//删除gg缓存
	_ = os.RemoveAll("./file/gg")
	//清空每日福地缓存
	table.CleanRoll()
	//刷新qq群信息
	refreshQQGroupInfo()
	//刷新玩家信息
	refreshPlayerInfo()
}

func refreshQQGroupInfo() {
	var qqs []int64
	infos := robotService.GetGroupInfo()
	msg := ""
	for _, qq := range infos {
		qqs = append(qqs, qq.UserId)
		if time.Now().Unix() > time.Unix(qq.JoinTime, 0).AddDate(0, 0, 7).Unix() {
			players := table.GetMSPlayerListByQQ(qq.UserId)
			if qq.Role == "member" && qq.UserId != 936831989 && (players == nil || len(players) == 0) {
				msg += fmt.Sprintf("QQ：%d  昵称：%s  加群时间：%s\n", qq.UserId, qq.Nickname, time.Unix(qq.JoinTime, 0).Format("2006-01-02 15:04:05"))
				if err := utils.POST("http://host:5700/set_group_kick", struct {
					UserId  int64 `json:"user_id"`
					GroupId int64 `json:"group_id"`
				}{
					UserId:  qq.UserId,
					GroupId: 769964102,
				}, nil); err != nil {
					log.Println(err)
				}
			}
		}
	}
	if msg != "" {
		msg = "以下成员超过7天未绑定角色:\n" + msg + "已被移除群！"
		robotService.SendGroupMsg(model.CQMessage{GroupId: 769964102}, msg)
	}

	table.CleanPlayer(qqs)

}

func refreshPlayerInfo() {
	playerList := table.GetMSPlayerList()
	errors := ""
	wg := sync.WaitGroup{}
	for _, player := range playerList {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			gg, err := robotService.GetGGData(name)
			if err != nil {
				errors += fmt.Sprintf("gg查询失败：【%s】\n", name)
			} else {
				if err = table.SaveMSPlayer(table.MSPlayer{
					Name:     gg.CharacterData.Name,
					Class:    gg.CharacterData.Class,
					Level:    gg.CharacterData.Level,
					Img:      gg.CharacterData.CharacterImageURL,
					IsMain:   gg.CharacterData.AchievementPoints != 0,
					Datatime: time.Now(),
				}); err != nil {
					errors += err.Error() + "\n"
				}
			}
		}(player.Name)
	}
	wg.Wait()
	if errors != "" {
		robotService.SendGroupMsg(model.CQMessage{GroupId: 769964102}, fmt.Sprintf("[刷新玩家信息异常]--->%s", errors))
	}
}
