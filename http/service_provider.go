package http

import (
	consolecontract "github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/facades"
	"github.com/chenyuIT/framework/http/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.RateLimiter = NewRateLimiter()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
