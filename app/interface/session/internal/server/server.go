package server

import (
	"flag"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/session/internal/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/session/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/interface/session/internal/server/grpc"
)

var configFile = flag.String("f", "etc/session.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	s.grpcSrv = grpc.New(c.RpcServerConf, service.New(c))

	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
