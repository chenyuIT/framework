package facades

import (
	"github.com/chenyuIT/framework/contracts/hash"
)

func Hash() hash.Hash {
	return App().MakeHash()
}
