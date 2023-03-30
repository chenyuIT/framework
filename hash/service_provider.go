package hash

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (hash *ServiceProvider) Register() {
	facades.Hash = NewApplication()
}

func (hash *ServiceProvider) Boot() {

}
