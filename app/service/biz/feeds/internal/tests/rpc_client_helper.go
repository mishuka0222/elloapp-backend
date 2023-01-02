package tests

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	feeds_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/client"
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
