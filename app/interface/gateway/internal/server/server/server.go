package server

import (
	"flag"
	"syscall"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/gateway/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/gateway/internal/server"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/gateway/internal/server/grpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	configFile = flag.String("f", "etc/gateway.yaml", "the config file")
)

type Server struct {
	grpcSrv *zrpc.RpcServer
	server  *server.Server
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)

	s.server = server.New(c)
	s.grpcSrv = grpc.New(c.RpcServerConf, s.server)

	go func() {
		s.grpcSrv.Start()
	}()

	return nil
}

func (s *Server) RunLoop() {
	if err := s.server.Serve(); err != nil {
		logx.Errorf("run server error: %v, quit...", err)
		commands.GSignal <- syscall.SIGQUIT
	}
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
	s.server.Close()
}
