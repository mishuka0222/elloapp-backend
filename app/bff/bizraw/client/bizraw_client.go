package bizraw_client

import (
	"context"
	"github.com/teamgram/proto/mtproto"
	"github.com/zeromicro/go-zero/zrpc"
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
