package service

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/gateway/gateway"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

// eGateOptions comet options.
type gatewayOptions struct {
	RoutineSize uint64
	RoutineChan uint64
}

// Gateway eGateClient is a gateway.
type Gateway struct {
	serverId string
	client   gateway_client.GatewayClient
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewSession new a comet.
func NewGateway(c zrpc.RpcClientConf) (*Gateway, error) {
	g := &Gateway{
		serverId: c.Endpoints[0],
		// client: make([]chan interface{}, options.RoutineSize),
		// options:     options,
	}

	cli, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("watchComet NewClient(%+v) error(%v)", c, err)
		return nil, err
	}
	g.client = gateway_client.NewGatewayClient(cli)
	g.ctx, g.cancel = context.WithCancel(context.Background())

	return g, nil
}

func (c *Gateway) Close() (err error) {
	c.cancel()
	return
}

// SendDataToGate
// egate.sendDataToGateway auth_key_id:long session_id:long payload:bytes = Bool;
func (c *Gateway) SendDataToGate(ctx context.Context, authKeyId, sessionId int64, payload []byte) (b bool, err error) {
	var (
		res *mtproto.Bool
	)

	res, err = c.client.GatewaySendDataToGateway(context.Background(), &gateway.TLGatewaySendDataToGateway{
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		Payload:   payload,
	})

	if err != nil {
		logx.Errorf("sendDataToGate error: %v", err)
		b = false
		return
	}

	b = mtproto.FromBool(res)
	return
}

//// NewGateway
//// NewComet new a comet.
//func NewGateway(data *naming.Instance, conf *Config, options gatewayOptions) (*Gateway, error) {
//	c := &Gateway{
//		serverID: data.Hostname,
//	}
//	var grpcAddr string
//	for _, addrs := range data.Addrs {
//		u, err := url.Parse(addrs)
//		if err == nil && u.Scheme == "grpc" {
//			grpcAddr = u.Host
//		}
//	}
//	if grpcAddr == "" {
//		return nil, fmt.Errorf("invalid grpc address:%v", data.Addrs)
//	}
//	var err error
//	if c.client, err = gateway_client.NewClient(grpcAddr, conf.WardenClient); err != nil {
//		return nil, err
//	}
//	return c, nil
//}
