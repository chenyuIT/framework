package facades

import (
	"github.com/chenyuIT/framework/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
