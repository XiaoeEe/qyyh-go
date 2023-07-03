package partyService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
	"qyyh-go/model"
	robotService "qyyh-go/service/robotSerice"
	"strings"
)

func GetPartyList(parm model.GetPartyListParm, _ *gin.Context) (data any, err error) {
	partys := table.GetPartyListByType(parm.Type)
	var partyList []model.GetPartyListReq
	for _, party := range partys {
		var (
			member []model.PartyPlayer
			i      int
		)
		for _, s := range strings.Split(party.Member, ",") {
			member = append(member, getPartyPlayer(s))
			i++
		}
		for ; i < party.Count; i++ {
			member = append(member, model.PartyPlayer{Name: ""})
		}
		partyList = append(partyList, model.GetPartyListReq{
			ID:       party.ID,
			Title:    party.Boss,
			Type:     party.Type,
			Remarks:  party.Remarks,
			Level:    party.Level,
			Time:     party.Time,
			Ch:       party.Channel,
			Point:    party.Point,
			Leader:   getPartyPlayer(party.Leader),
			Member:   member,
			CreateBy: party.CreateBy,
		})
	}

	return partyList, nil
}

func CreateParty(parm model.CreatePartyParm, c *gin.Context) (data any, err error) {
	var user table.User
	if u, ok := c.Get("user"); ok {
		user = u.(table.User)
	}
	err = table.CreateParty(table.Party{
		Leader:   parm.Leader,
		Boss:     strings.Join(parm.Boss, ","),
		Count:    parm.Count,
		Level:    parm.Level,
		Point:    parm.Point,
		Remarks:  parm.Remarks,
		Time:     parm.Time,
		Channel:  parm.Channel,
		CreateBy: user.Id,
		Type:     parm.Type,
	})
	BossList := map[string]string{
		"Lotus":                "斯乌",
		"Damien":               "戴米安",
		"Lucid":                "路西德",
		"Will":                 "威尔",
		"Guardian_Angel_Slime": "天使绿水灵王",
		"Gloom":                "至暗魔晶",
		"Verus_Hilla":          "觉醒希拉",
		"Darknell":             "亲卫队长敦凯尔",
		"Black_Mage":           "黑魔法师",
		"Chosen_Seren":         "神选者塞伦",
		"Kalos_the_Guardian":   "卡洛斯",
	}
	var chBoss []string
	for _, boss := range parm.Boss {
		chBoss = append(chBoss, BossList[boss])
	}
	if err == nil {
		s := fmt.Sprintf("有车头【%s】创建了队伍【%s】！\n发车时间：【%s】 发车线路：【%d】\n详情查看：https://qyyh.net/maplestory/party(备用网址：http://ah.qyyh.net:8888/maplestory/party)", parm.Leader, strings.Join(chBoss, ","), parm.Time, parm.Channel)
		robotService.SendGroupMsg(model.CQMessage{GroupId: 769964102}, s)
	}
	return err == nil, err
}

func DelParty(parm model.DelPartyParm, _ *gin.Context) (data any, err error) {
	err = table.DelParty(parm.ID)
	return err == nil, err
}

func JoinParty(parm model.JoinPartyParm, _ *gin.Context) (data any, err error) {
	party := table.GetPartyById(parm.ID)
	me := getPartyPlayer(parm.Me)
	if me.Point < party.Point {
		return nil, errors.New(fmt.Sprintf("宁未达到要求分数【%d】宁的分数【%d】，赶紧爬去跑旗！", party.Point, me.Point))
	}
	leader := getPartyPlayer(party.Leader)
	if me.QQ == leader.QQ {
		return nil, errors.New("宁上宁🐎的自己车呢！")
	}
	lastMember := model.PartyPlayer{}
	if mainJionPartyCheck(party.Boss, me.Name) {
		return nil, errors.New("路威及以上车只允许主号上车！")
	}
	if party.Member != "" {
		members := strings.Split(party.Member, ",")
		for _, memberName := range members {
			member := getPartyPlayer(memberName)
			if member.QQ == me.QQ {
				return nil, errors.New("宁都上车了，害搁着上宁🐎的车呢！")
			}
			if member.Point-member.UsePoint <= lastMember.Point-lastMember.UsePoint {
				lastMember = member
			}
		}
		if len(members) < party.Count {
			party.Member = party.Member + "," + me.Name
		} else {
			if me.Point-me.UsePoint-party.Point <= lastMember.Point-lastMember.UsePoint {
				return nil, errors.New("宁权重分比谁都低，上宁🐎的车呢！")
			}
			party.Member = strings.Replace(party.Member, lastMember.Name, me.Name, 1)
		}
	} else {
		party.Member = me.Name
	}
	err = table.UpdateParty(party)
	return err == nil, err
}

func LeaveParty(parm model.LeavePartyParm, c *gin.Context) (data any, err error) {
	var user table.User
	if u, ok := c.Get("user"); ok {
		user = u.(table.User)
	}
	party := table.GetPartyById(parm.ID)
	if getPartyPlayer(parm.Name).QQ != user.QQ {
		return nil, errors.New("不是你的角色。")
	}
	members := strings.Split(party.Member, ",")
	var newMembers []string
	for _, old := range members {
		if old != parm.Name {
			newMembers = append(newMembers, old)
		}
	}
	party.Member = strings.Join(newMembers, ",")
	err = table.UpdateParty(party)
	return err == nil, err
}

func getPartyPlayer(name string) model.PartyPlayer {
	if name == "" {
		return model.PartyPlayer{}
	}
	player := table.GetMSPlayer(name)
	return model.PartyPlayer{
		QQ:       player.QQ,
		Name:     player.Name,
		Level:    player.Level,
		Job:      player.Class,
		Point:    table.GetFlagRaceByQQ(player.QQ),
		UsePoint: table.GetUsePointByQQ(player.QQ),
		Img:      player.Img,
	}
}

func mainJionPartyCheck(boss, name string) bool {
	player := table.GetMSPlayer(name)
	if player.IsMain == true {
		return false
	}
	if table.GetPartyCountByQQ(player.QQ) >= int64(len(table.GetMPSDateList())) {
		return false
	}
	for _, b := range strings.Split(boss, ",") {
		if b != "Lotus" && b != "Damien" && b != "Guardian_Angel_Slime" {
			return true
		}
	}
	return false
}
