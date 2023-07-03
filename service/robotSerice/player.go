package robotService

import (
	"fmt"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"strings"
	"time"
)

func sendWelcomeMsg(message model.CQMessage) {
	group := table.GetGroupByGroupId(message.GroupId)
	if group.Group != 0 && group.WelcomeMsg != "" {
		SendGroupMsg(message, fmt.Sprintf("[CQ:at,qq=%d]%s", message.UserId, group.WelcomeMsg))
	}
}

func myBind(message model.CQMessage) {
	players := table.GetMSPlayerListByQQ(message.UserId)

	if players == nil || len(players) == 0 {
		SendGroupMsg(message, "宁未绑定任何角色")
	} else {
		var names []string
		for _, player := range players {
			names = append(names, player.Name)
		}
		SendGroupMsg(message, fmt.Sprintf("你绑定了：%s", strings.Join(names, ", ")))
	}
}

func bindPlayer(message model.CQMessage, ids string) {
	msidList := strings.Split(strings.TrimSpace(ids), ",")
	var ggs []model.GG
	for _, id := range msidList {
		gg, err := GetGGData(id)
		if err != nil || gg.CharacterData.Name == "" {
			SendGroupMsg(message, fmt.Sprintf("未查询到玩家%s, 请确认游戏ID", id))
			return
		} else {
			player := table.GetMSPlayer(id)
			if player.Name != "" {
				SendGroupMsg(message, fmt.Sprintf("角色【%s】已经被【QQ：%d】绑定", player.Name, player.QQ))
				return
			}
			ggs = append(ggs, gg)
		}
	}
	var err error
	for _, gg := range ggs {
		err = table.SaveMSPlayer(table.MSPlayer{
			Name:     gg.CharacterData.Name,
			Class:    gg.CharacterData.Class,
			Level:    gg.CharacterData.Level,
			Img:      gg.CharacterData.CharacterImageURL,
			QQ:       message.UserId,
			IsMain:   gg.CharacterData.LegionLevel != 0,
			Datatime: time.Now(),
		})
	}
	if err != nil {
		SendGroupMsg(message, err.Error())
	} else {
		SendGroupMsg(message, "绑定成功")
	}
}

func unBindPlayer(message model.CQMessage, ids string) {
	msidList := strings.Split(strings.TrimSpace(ids), ",")
	var players []table.MSPlayer
	for _, id := range msidList {
		player := table.GetMSPlayer(id)
		if player.Name != "" {
			if player.QQ != message.UserId {
				SendGroupMsg(message, fmt.Sprintf("【%s】不是你的号，你搁这解绑啥呢？？？", id))
				return
			}
			players = append(players, player)

		} else {
			SendGroupMsg(message, fmt.Sprintf("【%s】没绑定，你搁这解绑啥呢？？？", id))
		}
	}
	var err error
	for _, player := range players {
		err = table.DeleteMSPlayer(player)
	}
	if err != nil {
		SendGroupMsg(message, err.Error())
	} else {
		SendGroupMsg(message, "解绑成功")
	}
}
