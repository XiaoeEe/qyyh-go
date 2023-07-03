package userService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
	"qyyh-go/model"
	"qyyh-go/utils"
)

func QQLogin(parm model.QQParm, c *gin.Context) (data any, err error) {
	unionid := getUnionid(parm.AccessToken)
	user := table.GetUserByUnionid(unionid)
	if user.Id != 0 {
		return Login(user, c)
	}
	return
}

func QQRegedit(parm model.QQRegeditParm, c *gin.Context) (data any, err error) {
	uid := getUnionid(parm.AccessToken)
	if uid == "" {
		return nil, errors.New("获取Unionid失败")
	}
	if table.HasUserByUsername(parm.Username) {
		return nil, errors.New("用户名被占用了，换一个吧。")
	}
	user := table.User{
		Username: parm.Username,
		Password: parm.Password,
		Unionid:  uid,
	}
	data, err = table.CreatUser(user)
	if err == nil {
		return Login(data.(table.User), c)
	}
	return
}

func GetQQInfo(parm model.QQParm, _ *gin.Context) (data any, err error) {
	res := struct {
		Nickname string `json:"nickname"`
		Msg      string `json:"msg"`
	}{}
	if err = utils.Get(fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=101515771&openid=%s", parm.AccessToken, parm.OpenId), nil, &res); err != nil {
		return
	}
	if res.Nickname == "" {
		return nil, errors.New(res.Msg)
	}
	data = res
	return
}

func getUnionid(accessToken string) string {
	unionid := struct {
		Unionid string `json:"unionid"`
	}{}
	if err := utils.Get(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s&unionid=101515771&fmt=json", accessToken), nil, &unionid); err != nil {
		return ""
	}
	return unionid.Unionid
}
