package facades

import (
	foundationcontract "github.com/chenyuIT/framework/contracts/foundation"
	"github.com/chenyuIT/framework/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
