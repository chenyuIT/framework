package mock

import (
	accessmock "github.com/chenyuIT/framework/contracts/auth/access/mocks"
	authmock "github.com/chenyuIT/framework/contracts/auth/mocks"
	cachemock "github.com/chenyuIT/framework/contracts/cache/mocks"
	configmock "github.com/chenyuIT/framework/contracts/config/mocks"
	consolemock "github.com/chenyuIT/framework/contracts/console/mocks"
	ormmock "github.com/chenyuIT/framework/contracts/database/orm/mocks"
	eventmock "github.com/chenyuIT/framework/contracts/event/mocks"
	filesystemmock "github.com/chenyuIT/framework/contracts/filesystem/mocks"
	foundationmock "github.com/chenyuIT/framework/contracts/foundation/mocks"
	grpcmock "github.com/chenyuIT/framework/contracts/grpc/mocks"
	mailmock "github.com/chenyuIT/framework/contracts/mail/mocks"
	queuemock "github.com/chenyuIT/framework/contracts/queue/mocks"
	validatemock "github.com/chenyuIT/framework/contracts/validation/mocks"
	"github.com/chenyuIT/framework/foundation"
)

var app *foundationmock.Application

func App() *foundationmock.Application {
	if app == nil {
		app = &foundationmock.Application{}
		foundation.App = app
	}

	return app
}

func Artisan() *consolemock.Artisan {
	mockArtisan := &consolemock.Artisan{}
	App().On("MakeArtisan").Return(mockArtisan)

	return mockArtisan
}

func Auth() *authmock.Auth {
	mockAuth := &authmock.Auth{}
	App().On("MakeAuth").Return(mockAuth)

	return mockAuth
}

func Cache() (*cachemock.Cache, *cachemock.Driver, *cachemock.Lock) {
	mockCache := &cachemock.Cache{}
	App().On("MakeCache").Return(mockCache)

	return mockCache, &cachemock.Driver{}, &cachemock.Lock{}
}

func Config() *configmock.Config {
	mockConfig := &configmock.Config{}
	App().On("MakeConfig").Return(mockConfig)

	return mockConfig
}

func Event() (*eventmock.Instance, *eventmock.Task) {
	mockEvent := &eventmock.Instance{}
	App().On("MakeEvent").Return(mockEvent)

	return mockEvent, &eventmock.Task{}
}

func Gate() *accessmock.Gate {
	mockGate := &accessmock.Gate{}
	App().On("MakeGate").Return(mockGate)

	return mockGate
}

func Grpc() *grpcmock.Grpc {
	mockGrpc := &grpcmock.Grpc{}
	App().On("MakeGrpc").Return(mockGrpc)

	return mockGrpc
}

func Log() {
	App().On("MakeLog").Return(NewTestLog())
}

func Mail() *mailmock.Mail {
	mockMail := &mailmock.Mail{}
	App().On("MakeMail").Return(mockMail)

	return mockMail
}

func Orm() (*ormmock.Orm, *ormmock.Query, *ormmock.Transaction, *ormmock.Association) {
	mockOrm := &ormmock.Orm{}
	App().On("MakeOrm").Return(mockOrm)

	return mockOrm, &ormmock.Query{}, &ormmock.Transaction{}, &ormmock.Association{}
}

func Queue() (*queuemock.Queue, *queuemock.Task) {
	mockQueue := &queuemock.Queue{}
	App().On("MakeQueue").Return(mockQueue)

	return mockQueue, &queuemock.Task{}
}

func Storage() (*filesystemmock.Storage, *filesystemmock.Driver, *filesystemmock.File) {
	mockStorage := &filesystemmock.Storage{}
	mockDriver := &filesystemmock.Driver{}
	mockFile := &filesystemmock.File{}
	App().On("MakeStorage").Return(mockStorage)

	return mockStorage, mockDriver, mockFile
}

func Validation() (*validatemock.Validation, *validatemock.Validator, *validatemock.Errors) {
	mockValidation := &validatemock.Validation{}
	mockValidator := &validatemock.Validator{}
	mockErrors := &validatemock.Errors{}
	App().On("MakeValidation").Return(mockValidation)

	return mockValidation, mockValidator, mockErrors
}
