package grpc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		phonecall.RegisterRPCPhoneCallServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
