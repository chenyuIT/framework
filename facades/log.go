package facades

import (
	"github.com/chenyuIT/framework/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
