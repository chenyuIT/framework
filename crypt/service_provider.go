package crypt

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (crypt *ServiceProvider) Register() {
	facades.Crypt = NewApplication()
}

func (crypt *ServiceProvider) Boot() {

}
