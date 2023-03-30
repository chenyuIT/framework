package auth

import (
	"context"

	"github.com/chenyuIT/framework/auth/access"
	"github.com/chenyuIT/framework/auth/console"
	contractconsole "github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Auth = NewAuth(facades.Config.GetString("auth.defaults.guard"))
	facades.Gate = access.NewGate(context.Background())
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]contractconsole.Command{
		&console.JwtSecretCommand{},
		&console.PolicyMakeCommand{},
	})
}
