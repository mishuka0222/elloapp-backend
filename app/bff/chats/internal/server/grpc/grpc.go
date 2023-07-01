package grpc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/chats/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/chats/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		mtproto.RegisterRPCChatsServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
