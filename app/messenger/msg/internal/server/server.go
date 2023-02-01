package server

import (
	"flag"

	inbox_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/internal/config"
	msg_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
	mq      *kafka.ConsumerGroup
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// msg_helper
		msg.RegisterRPCMsgServer(
			grpcServer,
			msg_helper.New(
				msg_helper.Config{
					RpcServerConf:   c.RpcServerConf,
					Mysql:           c.Mysql,
					KV:              c.KV,
					IdgenClient:     c.IdgenClient,
					UserClient:      c.BizServiceClient,
					ChatClient:      c.BizServiceClient,
					ChannelsClient:  c.BizServiceClient,
					SyncClient:      c.SyncClient,
					InboxClient:     c.InboxClient,
					DialogClient:    c.BizServiceClient,
					MessageSharding: c.MessageSharding,
				}, nil))
	})

	go func() {
		s.grpcSrv.Start()
	}()

	s.mq = inbox_helper.New(inbox_helper.Config{
		RpcServerConf:   c.RpcServerConf,
		InboxConsumer:   c.InboxConsumer,
		Mysql:           c.Mysql,
		KV:              c.KV,
		IdgenClient:     c.IdgenClient,
		UserClient:      c.BizServiceClient,
		ChatClient:      c.BizServiceClient,
		ChannelsClient:  c.BizServiceClient,
		SyncClient:      c.SyncClient,
		BotSyncClient:   c.BotSyncClient,
		DialogClient:    c.BizServiceClient,
		MessageSharding: c.MessageSharding,
	})

	go func() {
		s.mq.Start()
	}()

	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
