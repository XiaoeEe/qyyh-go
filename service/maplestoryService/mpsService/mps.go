package mpsService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"qyyh-go/database/table"
	"qyyh-go/model"
)

func GetMPSDate(_ any, _ *gin.Context) (data any, err error) {
	data = table.GetMPSDateList()
	return
}

func GetMPS(parm model.GetMpsParm, _ *gin.Context) (data any, err error) {
	data = table.GetMPSByDate(parm.Date)
	return
}

func AddMPS(parm model.AddMPSParm, _ *gin.Context) (data any, err error) {
	if len(parm.Table) == 0 {
		return nil, errors.New("请输入要添加的内容")
	}
	oldData := table.GetMPSByDate(parm.Date)
	if len(oldData) != 0 {
		var ids []int64
		for _, mps := range parm.Table {
			mps.Date = parm.Date
			for _, old := range oldData {
				if mps.Name == old.Name {
					ids = append(ids, old.ID)
				}
			}
		}
		if len(ids) != 0 {
			if err = table.DelMPS(ids); err != nil {
				return
			}
		}
	}
	for i := range parm.Table {
		parm.Table[i].Date = parm.Date
	}
	err = table.CreateMPS(parm.Table)
	return
}

func GetMPSCount(parm model.GetMPSCountParm, _ *gin.Context) (data any, err error) {
	msps := table.GetMPSByDate(parm.Date)
	data = len(msps)
	return
}
