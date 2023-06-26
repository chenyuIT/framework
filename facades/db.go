package facades

import (
	"github.com/chenyuIT/framework/contracts/database/orm"
)

func Orm() orm.Orm {
	return App().MakeOrm()
}
