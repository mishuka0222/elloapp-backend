package tests

import (
	bizraw_client "github.com/teamgram/teamgram-server/app/service/biz/authorization/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewRPCClient() bizraw_client.RPCAuthorizationClient {
	return bizraw_client.NewAuthorizationClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "service.biz_service",
			InsecureSkipVerify: true,
		},
	}))
}
