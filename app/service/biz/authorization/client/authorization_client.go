package client

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/authorization"
	"github.com/zeromicro/go-zero/zrpc"
)

type RPCAuthorizationClient interface {
	AuthSingIN(ctx context.Context, in *authorization.AuthSingInRequest) (*authorization.AuthSingInRsp, error)
	AuthSingUP(ctx context.Context, in *authorization.AuthSingUpRequest) (*authorization.AuthSingUpRsp, error)
}

type defaultAuthorizationClient struct {
	cli zrpc.Client
}

func NewAuthorizationClient(cli zrpc.Client) RPCAuthorizationClient {
	return &defaultAuthorizationClient{
		cli: cli,
	}
}

func (c *defaultAuthorizationClient) AuthSingIN(ctx context.Context, in *authorization.AuthSingInRequest) (*authorization.AuthSingInRsp, error) {
	client := authorization.NewRPCAuthorizationClient(c.cli.Conn())
	return client.AuthSingIN(ctx, in)
}

func (c *defaultAuthorizationClient) AuthSingUP(ctx context.Context, in *authorization.AuthSingUpRequest) (*authorization.AuthSingUpRsp, error) {
	client := authorization.NewRPCAuthorizationClient(c.cli.Conn())
	return client.AuthSingUP(ctx, in)
}
