package sms

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (sms *ServiceProvider) Register() {
	facades.Sms = NewApplication()
}

func (sms *ServiceProvider) Boot() {

}
