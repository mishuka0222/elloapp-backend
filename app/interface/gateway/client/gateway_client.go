package gateway_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/gateway/gateway"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type GatewayClient interface {
	GatewaySendDataToGateway(ctx context.Context, in *gateway.TLGatewaySendDataToGateway) (*mtproto.Bool, error)
}

type defaultGatewayClient struct {
	cli zrpc.Client
}

func NewGatewayClient(cli zrpc.Client) GatewayClient {
	return &defaultGatewayClient{
		cli: cli,
	}
}

// GatewaySendDataToGateway
// gateway.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (m *defaultGatewayClient) GatewaySendDataToGateway(ctx context.Context, in *gateway.TLGatewaySendDataToGateway) (*mtproto.Bool, error) {
	client := gateway.NewRPCGatewayClient(m.cli.Conn())
	return client.GatewaySendDataToGateway(ctx, in)
}
