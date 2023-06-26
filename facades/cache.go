package facades

import (
	"github.com/chenyuIT/framework/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
