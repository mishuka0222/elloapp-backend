package tos_client

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type TosClient interface {
	HelpGetTermsOfServiceUpdate(ctx context.Context, in *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error)
	HelpAcceptTermsOfService(ctx context.Context, in *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error)
}

type defaultTosClient struct {
	cli zrpc.Client
}

func NewTosClient(cli zrpc.Client) TosClient {
	return &defaultTosClient{
		cli: cli,
	}
}

// HelpGetTermsOfServiceUpdate
// help.getTermsOfServiceUpdate#2ca51fd1 = help.TermsOfServiceUpdate;
func (m *defaultTosClient) HelpGetTermsOfServiceUpdate(ctx context.Context, in *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
	client := mtproto.NewRPCTosClient(m.cli.Conn())
	return client.HelpGetTermsOfServiceUpdate(ctx, in)
}

// HelpAcceptTermsOfService
// help.acceptTermsOfService#ee72f79a id:DataJSON = Bool;
func (m *defaultTosClient) HelpAcceptTermsOfService(ctx context.Context, in *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
	client := mtproto.NewRPCTosClient(m.cli.Conn())
	return client.HelpAcceptTermsOfService(ctx, in)
}
