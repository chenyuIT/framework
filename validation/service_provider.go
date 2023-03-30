package validation

import (
	consolecontract "github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/facades"
	"github.com/chenyuIT/framework/validation/console"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Validation = NewValidation()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
