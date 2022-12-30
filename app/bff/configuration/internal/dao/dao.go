package dao

import (
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/teamgram-server/app/bff/configuration/internal/config"
	configuration_client "github.com/teamgram/teamgram-server/app/service/biz/configuration/client"
)

type Dao struct {
	configuration_client.ConfigurationClient
}

func New(c config.Config) *Dao {
	return &Dao{
		ConfigurationClient: configuration_client.NewConfigurationClient(rpcx.GetCachedRpcClient(c.ConfigurationClient)),
	}
}
