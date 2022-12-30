package dao

import (
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/configuration/internal/config"
	configuration_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/configuration/client"
)

type Dao struct {
	configuration_client.ConfigurationClient
}

func New(c config.Config) *Dao {
	return &Dao{
		ConfigurationClient: configuration_client.NewConfigurationClient(rpcx.GetCachedRpcClient(c.ConfigurationClient)),
	}
}
