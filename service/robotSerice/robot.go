package robotService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"qyyh-go/model"
	"qyyh-go/utils"
	"strings"
)

func Robot(message model.CQMessage, _ *gin.Context) (data any, err error) {
	if message.PostType == "notice" && message.NoticeType == "group_increase" {
		sendWelcomeMsg(message)
	}

	if message.PostType == "message" && message.MessageType == "private" {
		m := strings.Split(strings.TrimSpace(message.Message), " ")
		switch strings.ToLower(m[0]) {
		case "注册账号":
			if len(m) == 3 {
				regedit(message, m[1], m[2])
			}
		case "绑定账号":
			if len(m) == 3 {
				bind(message, m[1], m[2])
			}
		}
	}

	if message.PostType == "message" && message.MessageType == "group" {
		m := strings.Split(strings.TrimSpace(message.Message), " ")
		switch strings.ToLower(m[0]) {
		case "机器人使用手册":
			if len(m) == 1 {
				SendGroupMsg(message, "https://qyyh.net/maplestory/robot/help")
			}
		case "绑定角色":
			if len(m) == 1 {
				go func(message model.CQMessage) {
					myBind(message)
				}(message)
			}
			if len(m) == 2 {
				go func(message model.CQMessage) {
					bindPlayer(message, m[1])
				}(message)
			}
		case "解绑角色":
			if len(m) == 2 {
				go func(message model.CQMessage) {
					unBindPlayer(message, m[1])
				}(message)
			}
		case "词条":
			if len(m) == 1 {
				getEntryList(message)
			}
			if len(m) == 2 {
				getEntry(message, m[1])
			}
			if len(m) >= 3 {
				setEntry(message, m)
			}
		case "删除词条":
			if len(m) == 2 {
				delEntry(message, m[1])
			}
		case "gg":
			if len(m) == 1 {
				go func(message model.CQMessage) {
					myGG(message)
				}(message)
			}
			if len(m) == 2 {
				go func(message model.CQMessage) {
					GGForName(message, m[1])
				}(message)
			}
		case "今日福地":
			if len(m) == 1 {
				rollHeavenly(message)
			}
		case "今天早上吃什么":
			if len(m) == 1 {
				rollMeal(message, 1)
			}
		case "今天中午吃什么":
			if len(m) == 1 {
				rollMeal(message, 2)
			}
		case "今天晚上吃什么":
			if len(m) == 1 {
				rollMeal(message, 3)
			}
		case "早餐登记":
			if len(m) == 2 {
				addMeal(message, m[1], 1)
			}
		case "午餐登记":
			if len(m) == 2 {
				addMeal(message, m[1], 2)
			}
		case "晚餐登记":
			if len(m) == 2 {
				addMeal(message, m[1], 3)
			}
		}
	}
	return
}

func SendGroupMsg(message model.CQMessage, str string) {
	msg := strings.TrimSpace(str)
	if message.MessageId != 0 {
		msg = fmt.Sprintf("[CQ:reply,id=%d]%s", message.MessageId, msg)
	}
	if err := utils.POST("http://host:5700/send_group_msg", model.SendGroupMessage{
		GroupId: message.GroupId,
		Message: msg,
	}, nil); err != nil {
		log.Println(err)
	}
}

func SendMessage(message model.CQMessage, str string) {
	if err := utils.POST("http://host:5700/send_private_msg", model.SendMessage{
		UserId:  message.UserId,
		GroupId: 769964102,
		Message: strings.TrimSpace(str),
	}, nil); err != nil {
		log.Println(err)
	}
}
