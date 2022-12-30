package tests

import (
	feeds_client "github.com/teamgram/teamgram-server/app/service/biz/feeds/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewRPCClient() feeds_client.FeedsClient {
	return feeds_client.NewFeedsClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "service.biz_service",
			InsecureSkipVerify: true,
		},
	}))
}
