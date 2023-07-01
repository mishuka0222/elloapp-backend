package tests

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	bizraw_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/client"
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
