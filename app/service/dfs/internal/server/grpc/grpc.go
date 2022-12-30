package grpc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		dfs.RegisterRPCDfsServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
