package schedule

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register() {
	facades.Schedule = NewApplication()
}

func (receiver *ServiceProvider) Boot() {

}
