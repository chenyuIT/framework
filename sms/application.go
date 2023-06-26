package sms

import (
	"github.com/chenyuIT/framework/contracts/sms"
)

const (
	VERSION         = "0.16 beta"
	ContentType     = "text/json"
	CacheSize   int = 64
	WriteBuffer int = 32
	Maxsendtime     = 60
)

type Application struct {
}

func NewApplication() sms.Sms {
	return NewSms()
}
