package facades

import (
	"github.com/chenyuIT/framework/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
