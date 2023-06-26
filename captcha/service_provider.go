package captcha

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (captcha *ServiceProvider) Register() {
	facades.Captcha = NewApplication()
}

func (captcha *ServiceProvider) Boot() {

}
