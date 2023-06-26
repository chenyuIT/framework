package cache

import (
	"github.com/chenyuIT/framework/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
