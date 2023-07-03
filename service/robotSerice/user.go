package robotService

import (
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/utils"
)

func regedit(message model.CQMessage, username, password string) {
	qqs := GetQQList()
	if !utils.Isin(qqs, message.UserId) {
		SendMessage(message, "你不在群里，不能使用快捷注册功能")
		return
	}
	if table.HasUserByUsername(username) {
		SendMessage(message, "用户名被占用了，换一个吧。")
		return
	}
	u := table.GetUserByQQ(message.UserId)
	if u.Id != 0 {
		SendMessage(message, "你的QQ已经注册过账号了。")
		return
	}
	if _, err := table.CreatUser(table.User{
		Username: username,
		Password: utils.MD5(password),
		QQ:       message.UserId,
	}); err != nil {
		SendMessage(message, err.Error())
	} else {
		SendMessage(message, "注册成功")
	}

}

func bind(message model.CQMessage, username, password string) {
	qqs := GetQQList()
	if !utils.Isin(qqs, message.UserId) {
		SendMessage(message, "你不在群里，不能使用快捷注册功能")
		return
	}
	if u := table.GetUserByQQ(message.UserId); u.Id != 0 {
		SendMessage(message, "你的QQ已经绑定过账号了。")
		return
	}
	u := table.GetUserByUsernameAndPassword(username, utils.MD5(password))
	if u.Id == 0 {
		SendMessage(message, "用户名或密码错误。")
		return
	}
	if u.QQ != 0 {
		SendMessage(message, "账号已经绑定过QQ了。")
		return
	}
	u.QQ = message.UserId
	if err := table.UpdateUser(u); err != nil {
		SendMessage(message, err.Error())
	} else {
		SendMessage(message, "绑定成功")
	}
}
