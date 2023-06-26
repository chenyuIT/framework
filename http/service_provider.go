package http

import (
	"github.com/chenyuIT/framework/contracts/cache"
	"github.com/chenyuIT/framework/contracts/config"
	consolecontract "github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/contracts/foundation"
	"github.com/chenyuIT/framework/contracts/http"
	"github.com/chenyuIT/framework/contracts/log"
	"github.com/chenyuIT/framework/contracts/validation"
	"github.com/chenyuIT/framework/http/console"
)

const Binding = "goravel.http"

var (
	ConfigFacade      config.Config
	CacheFacade       cache.Cache
	LogFacade         log.Log
	RateLimiterFacade http.RateLimiter
	ValidationFacade  validation.Validation
)

type ServiceProvider struct {
}

func (http *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewRateLimiter(), nil
	})
}

func (http *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	CacheFacade = app.MakeCache()
	LogFacade = app.MakeLog()
	RateLimiterFacade = app.MakeRateLimiter()
	ValidationFacade = app.MakeValidation()

	http.registerCommands(app)
}

func (http *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
