package task

import (
	"github.com/jasonlvhit/gocron"
)

func Init() {
	_ = gocron.Every(1).Days().At("08:00").Do(At8)
	//_ = gocron.Every(1).Days().At("00:00").Do(service.At0)

	gocron.Start()
}
