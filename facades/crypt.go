package facades

import (
	"github.com/chenyuIT/framework/contracts/crypt"
)

func Crypt() crypt.Crypt {
	return App().MakeCrypt()
}
