package filesystem

import (
	configcontract "github.com/chenyuIT/framework/contracts/config"
	filesystemcontract "github.com/chenyuIT/framework/contracts/filesystem"
	"github.com/chenyuIT/framework/contracts/foundation"
)

const Binding = "goravel.filesystem"

var ConfigFacade configcontract.Config
var StorageFacade filesystemcontract.Storage

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewStorage(app.MakeConfig()), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	StorageFacade = app.MakeStorage()
}
