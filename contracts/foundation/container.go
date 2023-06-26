package foundation

import (
	"github.com/chenyuIT/framework/contracts/auth"
	"github.com/chenyuIT/framework/contracts/auth/access"
	"github.com/chenyuIT/framework/contracts/cache"
	"github.com/chenyuIT/framework/contracts/config"
	"github.com/chenyuIT/framework/contracts/console"
	"github.com/chenyuIT/framework/contracts/crypt"
	"github.com/chenyuIT/framework/contracts/database/orm"
	"github.com/chenyuIT/framework/contracts/event"
	"github.com/chenyuIT/framework/contracts/filesystem"
	"github.com/chenyuIT/framework/contracts/grpc"
	"github.com/chenyuIT/framework/contracts/hash"
	"github.com/chenyuIT/framework/contracts/http"
	"github.com/chenyuIT/framework/contracts/log"
	"github.com/chenyuIT/framework/contracts/mail"
	"github.com/chenyuIT/framework/contracts/queue"
	"github.com/chenyuIT/framework/contracts/route"
	"github.com/chenyuIT/framework/contracts/schedule"
	"github.com/chenyuIT/framework/contracts/validation"
)

type Container interface {
	Bind(key any, callback func(app Application) (any, error))
	BindWith(key any, callback func(app Application, parameters map[string]any) (any, error))
	Instance(key, instance any)
	Make(key any) (any, error)
	MakeArtisan() console.Artisan
	MakeAuth() auth.Auth
	MakeCache() cache.Cache
	MakeConfig() config.Config
	MakeCrypt() crypt.Crypt
	MakeEvent() event.Instance
	MakeGate() access.Gate
	MakeGrpc() grpc.Grpc
	MakeHash() hash.Hash
	MakeLog() log.Log
	MakeMail() mail.Mail
	MakeOrm() orm.Orm
	MakeQueue() queue.Queue
	MakeRateLimiter() http.RateLimiter
	MakeRoute() route.Engine
	MakeSchedule() schedule.Schedule
	MakeStorage() filesystem.Storage
	MakeValidation() validation.Validation
	MakeWith(key any, parameters map[string]any) (any, error)
	Singleton(key any, callback func(app Application) (any, error))
}
