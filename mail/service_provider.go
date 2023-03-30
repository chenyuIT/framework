package mail

import (
	"github.com/chenyuIT/framework/contracts/queue"
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (route *ServiceProvider) Register() {
	facades.Mail = NewApplication()
}

func (route *ServiceProvider) Boot() {
	facades.Queue.Register([]queue.Job{
		&SendMailJob{},
	})
}
