package ocr

import (
	"github.com/chenyuIT/framework/facades"
)

type ServiceProvider struct {
}

func (ocr *ServiceProvider) Register() {
	facades.Ocr = NewApplication()
}

func (ocr *ServiceProvider) Boot() {

}
