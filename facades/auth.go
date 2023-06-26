package facades

import (
	"github.com/chenyuIT/framework/contracts/auth"
	"github.com/chenyuIT/framework/contracts/auth/access"
)

func Auth() auth.Auth {
	return App().MakeAuth()
}

func Gate() access.Gate {
	return App().MakeGate()
}
