package captcha

import (
	"github.com/chenyuIT/framework/contracts/captcha"
)

type Application struct {
}

func NewApplication() captcha.Captcha {
	return NewCaptcha()
}
