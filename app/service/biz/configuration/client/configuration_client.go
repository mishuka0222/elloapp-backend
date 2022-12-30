package configuration_client

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/configuration"
	"github.com/zeromicro/go-zero/zrpc"
)

type ConfigurationClient interface {
	ConfigurationGetCountryList(ctx context.Context, in *configuration.GetCountryListReq) (*configuration.GetCountryListResp, error)
}

type defaultConfigurationClient struct {
	cli zrpc.Client
}

func NewConfigurationClient(cli zrpc.Client) ConfigurationClient {
	return &defaultConfigurationClient{
		cli: cli,
	}
}

func (c *defaultConfigurationClient) ConfigurationGetCountryList(ctx context.Context, in *configuration.GetCountryListReq) (*configuration.GetCountryListResp, error) {
	client := configuration.NewRPCConfigurationClient(c.cli.Conn())
	return client.ConfigurationGetCountryList(ctx, in)
}
