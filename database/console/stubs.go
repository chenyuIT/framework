package console

type Stubs struct {
}

func (r Stubs) Model() string {
	return `package models

import (
	"github.com/chenyuIT/framework/database/orm"
)

type DummyModel struct {
	orm.Model
}
`
}
