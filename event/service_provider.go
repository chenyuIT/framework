package event

import (
	"github.com/chenyuIT/framework/contracts/console"
	eventConsole "github.com/chenyuIT/framework/event/console"
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Event = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console.Command{
		&eventConsole.EventMakeCommand{},
		&eventConsole.ListenerMakeCommand{},
	})
}
