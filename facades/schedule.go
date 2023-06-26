package facades

import (
	"github.com/chenyuIT/framework/contracts/schedule"
)

func Schedule() schedule.Schedule {
	return App().MakeSchedule()
}
