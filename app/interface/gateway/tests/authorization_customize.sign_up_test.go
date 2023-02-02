package tests

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/gateway/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/gateway/gateway"
	"testing"
)

func TestSignUp(t *testing.T) {

	gatewayClient := gateway_client.NewGatewayClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "interface.session",
		},
		Endpoints: []string{},
	}))

	hi, err := gatewayClient.GatewaySendDataToGateway(context.TODO(), &gateway.TLGatewaySendDataToGateway{
		Constructor: gateway.CRC32_UNKNOWN,
		AuthKeyId:   1937090286237253747,
		SessionId:   -6067022183794192689,
		Payload:     []byte{},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(hi)
}
