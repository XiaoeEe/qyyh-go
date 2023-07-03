package robotService

import (
	"fmt"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"regexp"
	"strconv"
)

func wantGollux(message model.CQMessage) {
	if err := table.CreateGollux(table.Gollux{QQ: message.UserId}); err != nil {
		SendGroupMsg(message, "失败: "+err.Error())
	} else {
		SendGroupMsg(message, "登记成功")
	}
}

func delGollux(message model.CQMessage) {
	if err := table.DelGollux(message.UserId); err != nil {
		SendGroupMsg(message, "失败: "+err.Error())
	} else {
		SendGroupMsg(message, "登记成功")
	}
}

func runGollux(message model.CQMessage, m []string) {
	var (
		ch    int
		value string
		err   error
	)
	for i := 2; i < len(m); i++ {
		value += m[i] + " "
	}
	sampleRegexp := regexp.MustCompile(`\d+`)
	if ch, err = strconv.Atoi(sampleRegexp.FindString(m[1])); err != nil {
		SendGroupMsg(message, "参数错误: "+err.Error())
		return
	}
	if ch < 1 || ch >= 30 {
		SendGroupMsg(message, "？？？")
		return
	}
	msg := "你这个车群里没人要啊~"
	golluxs := table.GetGolluxList()
	if len(golluxs) > 0 {
		msg = fmt.Sprintf("有三核准备在%d线发车啦！需要的小伙伴快上车啊！\n", ch)
		for _, gollux := range golluxs {
			msg += fmt.Sprintf("[CQ:at,qq=%d] ", gollux.QQ)
		}
		if value != "" {
			msg += fmt.Sprintf("\n司机留言【%s】", value)
		}
	}
	SendGroupMsg(message, msg)
}
