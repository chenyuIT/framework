package crypt

import (
	"github.com/chenyuIT/framework/contracts/crypt"
)

type Application struct {
}

func NewApplication() crypt.Crypt {
	return NewAES()
}
