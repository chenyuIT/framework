package facades

import "github.com/chenyuIT/framework/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
