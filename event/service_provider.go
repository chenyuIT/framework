package event

import (
	"github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/contracts/foundation"
	eventConsole "github.com/chenyuIT/framework/event/console"
)

const Binding = "goravel.event"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.MakeQueue()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]console.Command{
		&eventConsole.EventMakeCommand{},
		&eventConsole.ListenerMakeCommand{},
	})
}
