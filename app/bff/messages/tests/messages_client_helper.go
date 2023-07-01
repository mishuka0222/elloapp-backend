package tests

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	dialogs_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/messages/client"
)

func NewRPCClient() dialogs_client.MessagesClient {
	return dialogs_client.NewMessagesClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "bff.bff",
			InsecureSkipVerify: true,
		},
	}))
}
