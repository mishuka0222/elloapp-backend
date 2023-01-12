package client

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type RPCAuthorizationClient interface {
	AuthSignIn(ctx context.Context, in *authorization.AuthSignInRequest) (*mtproto.Auth_Authorization, error)
	AuthSignUp(ctx context.Context, in *authorization.AuthSignUpRequest) (*authorization.AuthSignUpRsp, error)
	Confirmation(ctx context.Context, in *authorization.ConfirmationRequest) (*authorization.ConfirmationResponse, error)
}

type defaultAuthorizationClient struct {
	cli zrpc.Client
}

func NewAuthorizationClient(cli zrpc.Client) RPCAuthorizationClient {
	return &defaultAuthorizationClient{
		cli: cli,
	}
}

func (c *defaultAuthorizationClient) AuthSignIn(ctx context.Context, in *authorization.AuthSignInRequest) (*mtproto.Auth_Authorization, error) {
	client := authorization.NewRPCAuthorizationClient(c.cli.Conn())
	return client.AuthSignIn(ctx, in)
}

func (c *defaultAuthorizationClient) AuthSignUp(ctx context.Context, in *authorization.AuthSignUpRequest) (*authorization.AuthSignUpRsp, error) {
	client := authorization.NewRPCAuthorizationClient(c.cli.Conn())
	return client.AuthSignUp(ctx, in)
}

func (c *defaultAuthorizationClient) Confirmation(ctx context.Context, in *authorization.ConfirmationRequest) (*authorization.ConfirmationResponse, error) {
	client := authorization.NewRPCAuthorizationClient(c.cli.Conn())
	return client.Confirmation(ctx, in)
}
