package bizraw_client

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"google.golang.org/grpc"
)

type OperationProxyServiceClient interface {
	mtproto.RPCBizClient
}

type defaultBizrawClient struct {
	cli zrpc.Client
}

func NewOperationProxyServiceClient(cli zrpc.Client) OperationProxyServiceClient {
	return &defaultBizrawClient{
		cli: cli,
	}
}

func (m defaultBizrawClient) BizInvokeBizDataRaw(ctx context.Context, in *mtproto.TLBizInvokeBizDataRaw, opts ...grpc.CallOption) (*mtproto.BizDataRaw, error) {
	client := mtproto.NewRPCBizClient(m.cli.Conn())
	return client.BizInvokeBizDataRaw(ctx, in, opts...)
}
