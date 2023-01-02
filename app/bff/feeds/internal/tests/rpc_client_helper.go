package tests

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	bizraw_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/client"
)

func NewRPCClient() bizraw_client.OperationProxyServiceClient {
	return bizraw_client.NewOperationProxyServiceClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "bff.bff",
			InsecureSkipVerify: true,
		},
	}))
}
