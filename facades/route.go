package facades

import (
	"github.com/chenyuIT/framework/contracts/route"
)

func Route() route.Engine {
	return App().MakeRoute()
}
