package mpsService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qyyh-go/model"
	robotService "qyyh-go/service/robotSerice"
	"sync"
)

func CheckName(parm model.CheckNameParm, _ *gin.Context) (data any, err error) {
	names := parm.Name
	s := ""
	wg := sync.WaitGroup{}
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			if _, err := robotService.GetGGData(name); err != nil {
				s += fmt.Sprintf("[%s] 角色名校验错误<br />", name)
			}
		}(name)
	}
	wg.Wait()
	data = s
	return
}
