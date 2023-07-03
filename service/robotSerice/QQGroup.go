package robotService

import (
	"log"
	"qyyh-go/model"
	"qyyh-go/utils"
)

func GetGroupInfo() []model.GroupUser {
	p := struct {
		GroupId int64 `json:"group_id"`
		NoCache bool  `json:"no_cache"`
	}{
		GroupId: 769964102,
		NoCache: true,
	}
	var res struct {
		Data    []model.GroupUser `json:"data"`
		Wording string            `json:"wording"`
	}
	err := utils.POST("http://host:5700/get_group_member_list", p, &res)
	if err != nil {
		log.Println(err)
		return nil
	}
	return res.Data
}

func GetQQList() []int64 {
	var qqs []int64
	for _, datum := range GetGroupInfo() {
		qqs = append(qqs, datum.UserId)
	}
	return qqs
}
