package facades

import (
	"github.com/chenyuIT/framework/contracts/queue"
)

func Queue() queue.Queue {
	return App().MakeQueue()
}
