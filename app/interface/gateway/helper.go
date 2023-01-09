package gateway_helper

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/gateway/internal/server/server"
)

func init() {
	zrpc.DontLogContentForMethod("/gateway.RPCGateway/GatewaySendDataToGateway")
}

type (
	Server = server.Server
)
