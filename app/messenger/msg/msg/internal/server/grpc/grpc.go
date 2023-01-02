package grpc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/internal/svc"
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		msgpb.RegisterRPCMsgServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
