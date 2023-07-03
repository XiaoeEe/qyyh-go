package robotService

import (
	"fmt"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/utils"
	"time"
)

func rollHeavenly(message model.CQMessage) {
	today := time.Now().Format("20060102")
	roll := table.GetRollByQQAndDate(message.UserId, today)
	msg := ""
	if roll.Id != 0 {
		msg = roll.Text
	} else {
		locations := table.GetLocations()
		l := utils.MicsSlice(locations, 4)
		n := utils.RandInt(5, 30)
		msg = fmt.Sprintf("上  星: %d线%s\n洗魔方: %d线%s\n洗火花: %d线%s\n开箱子: %d线%s\n打BOSS极大概率会在%d线出货", n[0]+1, l[0], n[1]+1, l[1], n[2]+1, l[2], n[3]+1, l[3], n[4]+1)
		if err := table.CreateRoll(table.Roll{
			QQ:   message.UserId,
			Date: today,
			Text: msg,
		}); err != nil {
			SendGroupMsg(message, err.Error())
			return
		}
	}
	SendGroupMsg(message, msg)
}

func addMeal(message model.CQMessage, name string, flag int) {
	meal := table.GetMealByNameAndFlag(name, flag)
	if meal.Id != 0 {
		SendGroupMsg(message, fmt.Sprintf("该菜品已由QQ号为%d的用户添加", meal.Userid))
		return
	}
	if err := table.CreateMeal(table.Meal{
		Name:   name,
		Flag:   flag,
		Userid: message.UserId,
	}); err != nil {
		SendGroupMsg(message, err.Error())
		return
	} else {
		SendGroupMsg(message, "添加成功")
	}
}

func rollMeal(message model.CQMessage, flag int) {
	SendGroupMsg(message, table.GetRollMeal(flag))
}
