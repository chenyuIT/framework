package facades

import "github.com/chenyuIT/framework/contracts/event"

func Event() event.Instance {
	return App().MakeEvent()
}
