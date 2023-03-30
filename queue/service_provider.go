package queue

import (
	"github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/facades"
	queueConsole "github.com/chenyuIT/framework/queue/console"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Queue = NewApplication()
}

func (receiver *ServiceProvider) Boot() {
	receiver.registerCommands()
}

func (receiver *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console.Command{
		&queueConsole.JobMakeCommand{},
	})
}
