package router

import (
	"github.com/gin-gonic/gin"
	"qyyh-go/service/maplestoryService/mpsService"
	"qyyh-go/service/maplestoryService/partyService"
	"qyyh-go/service/maplestoryService/playerService"
	"qyyh-go/service/robotSerice"
	"qyyh-go/service/userService"
)

var r *gin.Engine

func Init(e *gin.Engine) {
	r = e
	//robot
	router("/robot", robotService.Robot)
	//user
	router("/user/login", userService.Login)
	router("/user/regedit", userService.Regidit)
	router("/user/qqLogin", userService.QQLogin)
	router("/user/qqRegedit", userService.QQRegedit)
	router("/user/getQQInfo", userService.GetQQInfo)
	//mps
	router("/maplestory/ocr", mpsService.OCR)
	router("/maplestory/checkName", mpsService.CheckName)
	router("/maplestory/getMPSDate", mpsService.GetMPSDate)
	router("/maplestory/getMPS", mpsService.GetMPS)
	router("/maplestory/addMPS", mpsService.AddMPS)
	router("/maplestory/getMPSCount", mpsService.GetMPSCount)
	//player
	router("/maplestory/getInfo", playerService.GetMapleStoryInfo)
	//party
	router("/maplestory/party/getList", partyService.GetPartyList)
	router("/maplestory/party/create", partyService.CreateParty)
	router("/maplestory/party/delete", partyService.DelParty)
	router("/maplestory/party/join", partyService.JoinParty)
	router("/maplestory/party/leave", partyService.LeaveParty)

}

func router[T any](relativePath string, handler func(p T, c *gin.Context) (any, error)) {
	r.POST(relativePath, func(c *gin.Context) {
		var p T
		if err := c.ShouldBind(&p); err != nil {
			c.JSON(200, gin.H{"message": err.Error(), "code": 222})
			return
		}
		data, err := handler(p, c)
		if err != nil {
			c.JSON(200, gin.H{"message": err.Error(), "code": 777})
		} else if data != nil {
			c.JSON(200, gin.H{"message": "success", "code": 200, "data": data})
		}
	})
}
