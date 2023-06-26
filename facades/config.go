package facades

import (
	"github.com/chenyuIT/framework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
