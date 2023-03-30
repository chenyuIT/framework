package filesystem

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	facades.Storage = NewStorage()
}

func (database *ServiceProvider) Boot() {

}
