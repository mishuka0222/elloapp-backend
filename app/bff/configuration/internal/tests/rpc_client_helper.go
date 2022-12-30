package tests

import (
	configuration_client "github.com/teamgram/teamgram-server/app/bff/configuration/client"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewRPCClient() configuration_client.ConfigurationClient {
	return configuration_client.NewConfigurationClient(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "bff.bff",
			InsecureSkipVerify: true,
		},
	}))
}
