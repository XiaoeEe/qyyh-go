package userService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
	"qyyh-go/model"
)

func Login(parm table.User, c *gin.Context) (any, error) {
	ret := model.LoginRet{}
	user := table.GetUserByUsernameAndPassword(parm.Username, parm.Password)
	if user.Id == 0 {
		return nil, errors.New("用户名或密码错误。")
	}
	if token, err := table.CreateUserToken(user.Id); err != nil {
		return nil, err
	} else {
		ret.Id = user.Id
		ret.Username = user.Username
		ret.QQ = user.QQ
		c.SetCookie("token", token, 31536000, "/", "/", false, false)
	}
	return ret, nil
}

func Regidit(parm table.User, c *gin.Context) (data any, err error) {
	if table.HasUserByUsername(parm.Username) {
		return nil, errors.New("用户名被占用了，换一个吧。")
	}
	data, err = table.CreatUser(parm)
	if err == nil && data.(table.User).Id != 0 {
		return Login(data.(table.User), c)
	}
	return
}
