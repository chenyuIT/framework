package grpc

import (
	"github.com/chenyuIT/framework/contracts/foundation"
	"github.com/chenyuIT/framework/contracts/log"
)

const Binding = "goravel.grpc"

var LogFacade log.Log

type ServiceProvider struct {
}

func (route *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeConfig()), nil
	})
}

func (route *ServiceProvider) Boot(app foundation.Application) {
	LogFacade = app.MakeLog()
}
