package facades

import (
	"github.com/chenyuIT/framework/contracts/grpc"
)

func Grpc() grpc.Grpc {
	return App().MakeGrpc()
}
