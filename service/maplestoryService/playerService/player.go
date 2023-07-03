package playerService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"time"
)

func GetMapleStoryInfo(_ any, c *gin.Context) (data any, err error) {
	var user table.User
	if u, ok := c.Get("user"); ok {
		user = u.(table.User)
	}
	if user.QQ != 0 {
		playerList := table.GetMSPlayerListByQQ(user.QQ)
		if playerList == nil || len(playerList) == 0 {
			return nil, errors.New("你没有绑定QQ号，请通过私聊机器人的方式绑定QQ号。机器人使用手册：https://qyyh.net/maplestory/robot/help")
		}
		var (
			players     []model.MapleStoryPlayer
			subFlagrace int
		)
		for _, player := range playerList {
			p := GetUserInfo(player.Name)
			players = append(players, p)
			subFlagrace += p.Flagrace
		}
		data = model.MapleStoryInfo{
			Player:      players,
			SubFlagrace: subFlagrace,
		}
	}
	return
}

func GetUserInfo(name string) model.MapleStoryPlayer {
	year, week := time.Now().ISOWeek()
	mps := table.GetMPSByNameAndDate(name, fmt.Sprintf("%d-%d周", year, week-1))
	p := table.GetMSPlayer(name)
	return model.MapleStoryPlayer{
		Name:     p.Name,
		Class:    p.Class,
		Level:    p.Level,
		Img:      p.Img,
		DataTime: p.Datatime,
		Points:   mps.Points,
		Culvert:  mps.Culvert,
		Flagrace: mps.FlagRace,
	}
}
