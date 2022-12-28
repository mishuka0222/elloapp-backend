package tests

import (
	dialogs_client "github.com/teamgram/teamgram-server/app/bff/messages/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
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
