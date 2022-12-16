package tests

import (
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewDialogClient() dialog_client.DialogClient {
	return dialog_client.NewDialogClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "service.biz_service",
			InsecureSkipVerify: true,
		},
	}))
}
