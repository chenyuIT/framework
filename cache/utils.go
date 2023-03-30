package cache

import (
	"github.com/chenyuIT/framework/facades"
)

func prefix() string {
	return facades.Config.GetString("cache.prefix") + ":"
}
