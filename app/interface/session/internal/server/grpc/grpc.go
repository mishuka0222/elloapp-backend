package grpc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/session/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/session/session"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(c zrpc.RpcServerConf, service *service.Service) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		session.RegisterRPCSessionServer(grpcServer, service)
	})
	logx.Must(err)
	return s
}
